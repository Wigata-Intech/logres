package migrationx_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/shared/migrationx"
	"github.com/wigata-intech/logres/internal/shared/sqlite"
)

const (
	tblWidgets    = "widgets"
	createWidgets = "CREATE TABLE widgets (id INTEGER PRIMARY KEY);"
	dropWidgets   = "DROP TABLE widgets;"
)

type migFile struct {
	version  int64
	name     string
	up, down string
}

func openTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sqlite.Open(context.Background(), sqlite.Config{Path: ":memory:", MaxOpenConns: 1})
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })
	return db
}

// writeMigration writes an up/down pair; an empty side is skipped so tests can
// create orphaned files.
func writeMigration(t *testing.T, dir string, f migFile) {
	t.Helper()
	base := fmt.Sprintf("%d_%s", f.version, f.name)
	if f.up != "" {
		require.NoError(t, os.WriteFile(filepath.Join(dir, base+"_up.sql"), []byte(f.up), 0o600))
	}
	if f.down != "" {
		require.NoError(t, os.WriteFile(filepath.Join(dir, base+"_down.sql"), []byte(f.down), 0o600))
	}
}

func tableExists(ctx context.Context, t *testing.T, db *sql.DB, name string) bool {
	t.Helper()
	var got string
	err := db.QueryRowContext(ctx,
		"SELECT name FROM sqlite_master WHERE type='table' AND name=?", name).Scan(&got)
	if err == sql.ErrNoRows {
		return false
	}
	require.NoError(t, err)
	return got == name
}

// migratorCase is the shared shape for every Migrator scenario test: arrange
// files + options, run a scenario (which owns its own assertions), then match
// the terminal error.
type migratorCase struct {
	name    string
	files   []migFile
	opts    []migrationx.Option
	run     func(t *testing.T, ctx context.Context, m *migrationx.Migrator, db *sql.DB, dir string) error
	wantErr string
}

func runMigratorCases(t *testing.T, cases []migratorCase) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dir := t.TempDir()
			for _, f := range tc.files {
				writeMigration(t, dir, f)
			}
			db := openTestDB(t)
			m := migrationx.New(db, os.DirFS(dir), tc.opts...)

			err := tc.run(t, context.Background(), m, db, dir)
			if tc.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestUp(t *testing.T) {
	t.Parallel()
	runMigratorCases(t, []migratorCase{
		{
			name:  "applies pending then is idempotent",
			files: []migFile{{1, tblWidgets, createWidgets, dropWidgets}},
			run: func(t *testing.T, ctx context.Context, m *migrationx.Migrator, db *sql.DB, _ string) error {
				applied, err := m.Up(ctx)
				require.NoError(t, err)
				require.Len(t, applied, 1)
				assert.True(t, tableExists(ctx, t, db, tblWidgets))

				again, err := m.Up(ctx)
				require.NoError(t, err)
				assert.Empty(t, again)
				return nil
			},
		},
		{
			name:  "bad sql surfaces an error",
			files: []migFile{{1, "broken", "NOT VALID SQL;", dropWidgets}},
			run: func(_ *testing.T, ctx context.Context, m *migrationx.Migrator, _ *sql.DB, _ string) error {
				_, err := m.Up(ctx)
				return err
			},
			wantErr: "apply",
		},
		{
			name:  "missing up file is rejected",
			files: []migFile{{1, "orphan", "", dropWidgets}},
			run: func(_ *testing.T, ctx context.Context, m *migrationx.Migrator, _ *sql.DB, _ string) error {
				_, err := m.Up(ctx)
				return err
			},
			wantErr: "missing up file",
		},
	})
}

func TestDown(t *testing.T) {
	t.Parallel()
	runMigratorCases(t, []migratorCase{
		{
			name:  "reverts most recent then reports empty",
			files: []migFile{{1, tblWidgets, createWidgets, dropWidgets}},
			run: func(t *testing.T, ctx context.Context, m *migrationx.Migrator, db *sql.DB, _ string) error {
				_, err := m.Up(ctx)
				require.NoError(t, err)

				reverted, err := m.Down(ctx)
				require.NoError(t, err)
				assert.Equal(t, int64(1), reverted.Version)
				assert.False(t, tableExists(ctx, t, db, tblWidgets))

				_, err = m.Down(ctx)
				assert.ErrorIs(t, err, migrationx.ErrNoMigrations)
				return nil
			},
		},
		{
			name:  "applied version without a source file errors",
			files: []migFile{{1, tblWidgets, createWidgets, dropWidgets}},
			run: func(t *testing.T, ctx context.Context, m *migrationx.Migrator, _ *sql.DB, dir string) error {
				_, err := m.Up(ctx)
				require.NoError(t, err)
				require.NoError(t, os.Remove(filepath.Join(dir, "1_widgets_up.sql")))
				require.NoError(t, os.Remove(filepath.Join(dir, "1_widgets_down.sql")))
				_, err = m.Down(ctx)
				return err
			},
			wantErr: "no source file",
		},
	})
}

func TestStatus(t *testing.T) {
	t.Parallel()
	runMigratorCases(t, []migratorCase{
		{
			name: "reports pending before and applied after Up",
			files: []migFile{
				{10, "first", createWidgets, dropWidgets},
				{20, "second", "CREATE TABLE gadgets (id INTEGER PRIMARY KEY);", "DROP TABLE gadgets;"},
			},
			opts: []migrationx.Option{migrationx.WithTable("custom_migrations")},
			run: func(t *testing.T, ctx context.Context, m *migrationx.Migrator, _ *sql.DB, _ string) error {
				before, err := m.Status(ctx)
				require.NoError(t, err)
				require.Len(t, before, 2)
				assert.Equal(t, int64(10), before[0].Version) // ascending
				assert.False(t, before[0].Applied)

				_, err = m.Up(ctx)
				require.NoError(t, err)

				after, err := m.Status(ctx)
				require.NoError(t, err)
				for _, s := range after {
					assert.True(t, s.Applied, "version %d", s.Version)
					assert.False(t, s.AppliedAt.IsZero())
				}
				return nil
			},
		},
		{
			name:  "invalid table name falls back to default",
			files: []migFile{{1, tblWidgets, createWidgets, dropWidgets}},
			opts:  []migrationx.Option{migrationx.WithTable("bad-name;drop")},
			run: func(_ *testing.T, ctx context.Context, m *migrationx.Migrator, _ *sql.DB, _ string) error {
				_, err := m.Up(ctx)
				return err
			},
		},
	})
}

func TestCreate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		input      string
		wantErr    bool
		wantSlugIn string
	}

	cases := []testCase{
		{name: "sanitizes and writes pair", input: "Add Admin Users!", wantSlugIn: "add_admin_users"},
		{name: "empty after sanitize errors", input: "!!!", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dir := t.TempDir()

			up, down, err := migrationx.Create(dir, tc.input)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Contains(t, up, tc.wantSlugIn+"_up.sql")
			assert.Contains(t, down, tc.wantSlugIn+"_down.sql")
			assert.FileExists(t, up)
			assert.FileExists(t, down)

			body, err := os.ReadFile(up) //nolint:gosec // path from Create, test-controlled
			require.NoError(t, err)
			assert.True(t, strings.HasPrefix(string(body), "-- migrate: up"))
		})
	}
}
