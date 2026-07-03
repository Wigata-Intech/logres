// Package migrationx is a dependency-free, goose-style SQL migration runner.
//
// A migration is a pair of .sql files sharing a unix-timestamp version:
// {version}_{name}_up.sql and {version}_{name}_down.sql. Applied versions are
// recorded in a schema_migrations table so Up is idempotent. Each migration
// runs in its own transaction over plain database/sql.
package migrationx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const defaultTable = "schema_migrations"

const (
	suffixUp   = "_up"
	suffixDown = "_down"
)

// ErrNoMigrations is returned by Down when nothing has been applied.
var ErrNoMigrations = errors.New("migrationx: no applied migrations to revert")

var nameSanitizer = regexp.MustCompile(`[^a-z0-9]+`)

// identRE validates the table name, which is interpolated into SQL (a table
// name cannot be a bound parameter).
var identRE = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

type Migration struct {
	Version int64
	Name    string
	Up      string
	Down    string
}

type Status struct {
	Version   int64
	Name      string
	Applied   bool
	AppliedAt time.Time
}

type Migrator struct {
	db    *sql.DB
	fsys  fs.FS
	log   *slog.Logger
	table string
}

type Option func(*Migrator)

// WithTable overrides the bookkeeping table name. The name must be a valid SQL
// identifier; an invalid value is ignored and the default is kept.
func WithTable(name string) Option {
	return func(m *Migrator) {
		if identRE.MatchString(name) {
			m.table = name
		}
	}
}

func WithLogger(log *slog.Logger) Option {
	return func(m *Migrator) {
		if log != nil {
			m.log = log
		}
	}
}

// New builds a Migrator reading .sql files from fsys (e.g. os.DirFS("migration")).
func New(db *sql.DB, fsys fs.FS, opts ...Option) *Migrator {
	m := &Migrator{
		db:    db,
		fsys:  fsys,
		log:   slog.Default(),
		table: defaultTable,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Create writes an empty up/down migration pair into dir and returns their
// paths. The filename is `{unix}_{sanitized-name}_{up|down}.sql`.
func Create(dir, name string) (upPath, downPath string, err error) {
	slug := sanitizeName(name)
	if slug == "" {
		return "", "", errors.New("migrationx: migration name is empty after sanitizing")
	}

	if err = os.MkdirAll(dir, 0o750); err != nil {
		return "", "", fmt.Errorf("migrationx.Create: %w", err)
	}

	version := time.Now().Unix()
	base := strconv.FormatInt(version, 10) + "_" + slug
	upPath = filepath.Join(dir, base+suffixUp+".sql")
	downPath = filepath.Join(dir, base+suffixDown+".sql")

	files := []struct {
		path, body string
	}{
		{upPath, fmt.Sprintf("-- migrate: up\n-- %s\n\n", name)},
		{downPath, fmt.Sprintf("-- migrate: down\n-- %s\n\n", name)},
	}
	for _, f := range files {
		if err = os.WriteFile(f.path, []byte(f.body), 0o600); err != nil {
			return "", "", fmt.Errorf("migrationx.Create: %w", err)
		}
	}

	return upPath, downPath, nil
}

// Up applies every migration not yet recorded, oldest first, each in its own
// transaction. It returns the migrations it applied (possibly empty).
func (m *Migrator) Up(ctx context.Context) ([]Migration, error) {
	if err := m.ensureTable(ctx); err != nil {
		return nil, err
	}

	all, err := m.load()
	if err != nil {
		return nil, err
	}

	applied, err := m.appliedVersions(ctx)
	if err != nil {
		return nil, err
	}

	var done []Migration
	for _, mig := range all {
		if _, ok := applied[mig.Version]; ok {
			continue
		}
		if err := m.runUp(ctx, mig); err != nil {
			return done, err
		}
		m.log.InfoContext(ctx, "migration applied",
			slog.Int64("version", mig.Version), slog.String("name", mig.Name))
		done = append(done, mig)
	}

	return done, nil
}

// Down reverts the most recently applied migration in a single transaction.
// It returns ErrNoMigrations if nothing has been applied.
func (m *Migrator) Down(ctx context.Context) (*Migration, error) {
	if err := m.ensureTable(ctx); err != nil {
		return nil, err
	}

	all, err := m.load()
	if err != nil {
		return nil, err
	}

	applied, err := m.appliedVersions(ctx)
	if err != nil {
		return nil, err
	}
	if len(applied) == 0 {
		return nil, ErrNoMigrations
	}

	var latest int64 = -1
	for v := range applied {
		if v > latest {
			latest = v
		}
	}

	idx := -1
	for i := range all {
		if all[i].Version == latest {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, fmt.Errorf("migrationx.Down: applied version %d has no source file", latest)
	}

	mig := all[idx]
	if err := m.runDown(ctx, mig); err != nil {
		return nil, err
	}
	m.log.InfoContext(ctx, "migration reverted",
		slog.Int64("version", mig.Version), slog.String("name", mig.Name))

	return &mig, nil
}

// Status returns every known migration with its applied state, oldest first.
func (m *Migrator) Status(ctx context.Context) ([]Status, error) {
	if err := m.ensureTable(ctx); err != nil {
		return nil, err
	}

	all, err := m.load()
	if err != nil {
		return nil, err
	}

	appliedAt, err := m.appliedTimes(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]Status, 0, len(all))
	for _, mig := range all {
		at, ok := appliedAt[mig.Version]
		out = append(out, Status{
			Version:   mig.Version,
			Name:      mig.Name,
			Applied:   ok,
			AppliedAt: at,
		})
	}

	return out, nil
}

func (m *Migrator) runUp(ctx context.Context, mig Migration) error {
	return m.inTx(ctx, func(tx *sql.Tx) error {
		if _, err := tx.ExecContext(ctx, mig.Up); err != nil {
			return fmt.Errorf("apply %d_%s: %w", mig.Version, mig.Name, err)
		}
		// #nosec G202 -- m.table is a validated identifier, never user input.
		_, err := tx.ExecContext(ctx,
			"INSERT INTO "+m.table+" (version, name, applied_at) VALUES (?, ?, ?)",
			mig.Version, mig.Name, time.Now().UTC().Format(time.RFC3339Nano))
		if err != nil {
			return fmt.Errorf("record %d_%s: %w", mig.Version, mig.Name, err)
		}
		return nil
	})
}

func (m *Migrator) runDown(ctx context.Context, mig Migration) error {
	return m.inTx(ctx, func(tx *sql.Tx) error {
		if _, err := tx.ExecContext(ctx, mig.Down); err != nil {
			return fmt.Errorf("revert %d_%s: %w", mig.Version, mig.Name, err)
		}
		// #nosec G202 -- m.table is a validated identifier, never user input.
		_, err := tx.ExecContext(ctx,
			"DELETE FROM "+m.table+" WHERE version = ?", mig.Version)
		if err != nil {
			return fmt.Errorf("unrecord %d_%s: %w", mig.Version, mig.Name, err)
		}
		return nil
	})
}

// inTx runs fn inside a transaction, rolling back on error.
func (m *Migrator) inTx(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("migrationx: begin: %w", err)
	}
	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("migrationx: commit: %w", err)
	}
	return nil
}

func (m *Migrator) ensureTable(ctx context.Context) error {
	// #nosec G202 -- m.table is a validated identifier, never user input.
	q := "CREATE TABLE IF NOT EXISTS " + m.table + " (" +
		"version INTEGER PRIMARY KEY, " +
		"name TEXT NOT NULL, " +
		"applied_at TEXT NOT NULL)"
	if _, err := m.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("migrationx: ensure %s: %w", m.table, err)
	}
	return nil
}

func (m *Migrator) appliedVersions(ctx context.Context) (map[int64]struct{}, error) {
	times, err := m.appliedTimes(ctx)
	if err != nil {
		return nil, err
	}
	set := make(map[int64]struct{}, len(times))
	for v := range times {
		set[v] = struct{}{}
	}
	return set, nil
}

func (m *Migrator) appliedTimes(ctx context.Context) (map[int64]time.Time, error) {
	// #nosec G202 -- m.table is a validated identifier, never user input.
	rows, err := m.db.QueryContext(ctx, "SELECT version, applied_at FROM "+m.table)
	if err != nil {
		return nil, fmt.Errorf("migrationx: query %s: %w", m.table, err)
	}
	defer func() { _ = rows.Close() }()

	out := make(map[int64]time.Time)
	for rows.Next() {
		var (
			version int64
			raw     string
		)
		if err := rows.Scan(&version, &raw); err != nil {
			return nil, fmt.Errorf("migrationx: scan %s: %w", m.table, err)
		}
		at, _ := time.Parse(time.RFC3339Nano, raw)
		out[version] = at
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("migrationx: iterate %s: %w", m.table, err)
	}
	return out, nil
}

// load reads and pairs every migration file in the source FS, sorted ascending.
func (m *Migrator) load() ([]Migration, error) {
	entries, err := fs.ReadDir(m.fsys, ".")
	if err != nil {
		return nil, fmt.Errorf("migrationx: read source: %w", err)
	}

	byVersion := make(map[int64]*Migration)
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}

		version, name, dir, ok := parseFilename(e.Name())
		if !ok {
			continue
		}

		body, err := fs.ReadFile(m.fsys, e.Name())
		if err != nil {
			return nil, fmt.Errorf("migrationx: read %s: %w", e.Name(), err)
		}

		mig := byVersion[version]
		if mig == nil {
			mig = &Migration{Version: version, Name: name}
			byVersion[version] = mig
		}
		switch dir {
		case suffixUp:
			mig.Up = string(body)
		case suffixDown:
			mig.Down = string(body)
		}
	}

	out := make([]Migration, 0, len(byVersion))
	for _, mig := range byVersion {
		if mig.Up == "" {
			return nil, fmt.Errorf("migrationx: migration %d_%s missing up file", mig.Version, mig.Name)
		}
		out = append(out, *mig)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Version < out[j].Version })

	return out, nil
}

// parseFilename splits `{version}_{name}_{up|down}.sql` into its parts.
func parseFilename(filename string) (version int64, name, dir string, ok bool) {
	stem := strings.TrimSuffix(filename, ".sql")

	switch {
	case strings.HasSuffix(stem, suffixUp):
		dir = suffixUp
	case strings.HasSuffix(stem, suffixDown):
		dir = suffixDown
	default:
		return 0, "", "", false
	}
	stem = strings.TrimSuffix(stem, dir)

	prefix, rest, found := strings.Cut(stem, "_")
	if !found || rest == "" {
		return 0, "", "", false
	}
	v, err := strconv.ParseInt(prefix, 10, 64)
	if err != nil {
		return 0, "", "", false
	}

	return v, rest, dir, true
}

func sanitizeName(name string) string {
	s := nameSanitizer.ReplaceAllString(strings.ToLower(name), "_")
	return strings.Trim(s, "_")
}
