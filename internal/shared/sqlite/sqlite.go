// Package sqlite opens a configured *sql.DB backed by the pure-Go
// modernc.org/sqlite driver (no CGO). It is the default storage engine for
// Logres. Everything above it depends only on database/sql, so another engine
// can be supported later by adding an equivalent Open.
package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite" // registers the "sqlite" database/sql driver
)

const driverName = "sqlite"

// Config controls how the SQLite database is opened and pooled. Zero values
// fall back to the defaults noted below.
type Config struct {
	Path            string        // file path; ":memory:" for an in-memory database
	Shared          bool          // with in-memory Path, use one shared-cache DB (tests)
	BusyTimeout     time.Duration // lock wait before SQLITE_BUSY (default 5s)
	MaxOpenConns    int           // pool size; keep modest for a single file (default 16)
	MaxIdleConns    int           // idle pool size (default MaxOpenConns)
	ConnMaxIdleTime time.Duration // idle connection lifetime (default 5m)
}

func (c Config) withDefaults() Config {
	if c.BusyTimeout <= 0 {
		c.BusyTimeout = 5 * time.Second
	}
	if c.MaxOpenConns <= 0 {
		c.MaxOpenConns = 16
	}
	if c.MaxIdleConns <= 0 {
		c.MaxIdleConns = c.MaxOpenConns
	}
	if c.ConnMaxIdleTime <= 0 {
		c.ConnMaxIdleTime = 5 * time.Minute
	}
	return c
}

// dsn builds a DSN whose pragmas apply to every pooled connection: WAL +
// NORMAL synchronous for durable concurrency, foreign keys off by design.
func (c Config) dsn() string {
	path := c.Path
	if path == "" {
		path = ":memory:"
	}

	pragmas := []string{
		"busy_timeout(" + strconv.FormatInt(c.BusyTimeout.Milliseconds(), 10) + ")",
		"journal_mode(WAL)",
		"synchronous(NORMAL)",
		"foreign_keys(0)",
	}

	q := url.Values{}
	for _, p := range pragmas {
		q.Add("_pragma", p)
	}
	if c.Shared && strings.Contains(path, ":memory:") {
		q.Set("cache", "shared")
	}

	return "file:" + path + "?" + q.Encode()
}

// Open opens the database, applies pool limits, and pings it. The caller owns
// and must Close the returned *sql.DB.
func Open(ctx context.Context, cfg Config) (*sql.DB, error) {
	cfg = cfg.withDefaults()

	db, err := sql.Open(driverName, cfg.dsn())
	if err != nil {
		return nil, fmt.Errorf("sqlite.Open: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("sqlite.Open: ping: %w", err)
	}

	return db, nil
}
