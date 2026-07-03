package cmd

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/wigata-intech/logres/internal/shared/cli"
	"github.com/wigata-intech/logres/internal/shared/migrationx"
	"github.com/wigata-intech/logres/internal/shared/sqlite"
)

const (
	defaultDBPath       = "db/logres.db"
	defaultMigrationDir = "migration"
)

// migrateCommand groups the migration subcommands.
func migrateCommand() *cli.Command {
	c := cli.New("migrate", "create and apply database migrations")
	c.Sub(migrateCreateCommand())
	c.Sub(migrateUpCommand())
	c.Sub(migrateDownCommand())
	c.Sub(migrateStatusCommand())
	return c
}

func migrateCreateCommand() *cli.Command {
	c := cli.New("create", "generate a new up/down migration pair")
	fs := flag.NewFlagSet("migrate create", flag.ContinueOnError)
	dir := fs.String("dir", defaultMigrationDir, "migration directory")
	c.Flags = fs
	c.Exec = func(_ context.Context, args []string) error {
		if len(args) == 0 {
			return errors.New("migrate create: migration name is required")
		}
		up, down, err := migrationx.Create(*dir, strings.Join(args, " "))
		if err != nil {
			return err
		}
		fmt.Printf("created:\n  %s\n  %s\n", up, down)
		return nil
	}
	return c
}

func migrateUpCommand() *cli.Command {
	c := cli.New("up", "apply all pending migrations")
	dbPath, dir, fs := migrateDBFlags("migrate up")
	c.Flags = fs
	c.Exec = func(ctx context.Context, _ []string) error {
		m, db, err := openMigrator(ctx, *dbPath, *dir)
		if err != nil {
			return err
		}
		defer func() { _ = db.Close() }()

		applied, err := m.Up(ctx)
		if err != nil {
			return err
		}
		if len(applied) == 0 {
			fmt.Println("already up to date")
			return nil
		}
		for _, mig := range applied {
			fmt.Printf("applied  %d_%s\n", mig.Version, mig.Name)
		}
		return nil
	}
	return c
}

func migrateDownCommand() *cli.Command {
	c := cli.New("down", "revert the most recently applied migration")
	dbPath, dir, fs := migrateDBFlags("migrate down")
	c.Flags = fs
	c.Exec = func(ctx context.Context, _ []string) error {
		m, db, err := openMigrator(ctx, *dbPath, *dir)
		if err != nil {
			return err
		}
		defer func() { _ = db.Close() }()

		mig, err := m.Down(ctx)
		if errors.Is(err, migrationx.ErrNoMigrations) {
			fmt.Println("nothing to revert")
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("reverted %d_%s\n", mig.Version, mig.Name)
		return nil
	}
	return c
}

func migrateStatusCommand() *cli.Command {
	c := cli.New("status", "show applied and pending migrations")
	dbPath, dir, fs := migrateDBFlags("migrate status")
	c.Flags = fs
	c.Exec = func(ctx context.Context, _ []string) error {
		m, db, err := openMigrator(ctx, *dbPath, *dir)
		if err != nil {
			return err
		}
		defer func() { _ = db.Close() }()

		statuses, err := m.Status(ctx)
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "VERSION\tNAME\tSTATUS\tAPPLIED AT")
		for _, s := range statuses {
			state, at := "pending", "-"
			if s.Applied {
				state = "applied"
				at = s.AppliedAt.Format(time.RFC3339)
			}
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", s.Version, s.Name, state, at)
		}
		return w.Flush()
	}
	return c
}

// migrateDBFlags returns a flag set carrying the -db and -dir flags shared by
// the up/down/status subcommands.
func migrateDBFlags(name string) (dbPath, dir *string, fs *flag.FlagSet) {
	fs = flag.NewFlagSet(name, flag.ContinueOnError)
	dbPath = fs.String("db", defaultDBPath, "sqlite database path")
	dir = fs.String("dir", defaultMigrationDir, "migration directory")
	return dbPath, dir, fs
}

func openMigrator(ctx context.Context, dbPath, dir string) (*migrationx.Migrator, *sql.DB, error) {
	if parent := filepath.Dir(dbPath); parent != "." {
		if err := os.MkdirAll(parent, 0o750); err != nil {
			return nil, nil, fmt.Errorf("migrate: create db dir: %w", err)
		}
	}
	db, err := sqlite.Open(ctx, sqlite.Config{Path: dbPath})
	if err != nil {
		return nil, nil, err
	}
	return migrationx.New(db, os.DirFS(dir)), db, nil
}
