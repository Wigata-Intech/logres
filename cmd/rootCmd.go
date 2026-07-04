// Package cmd wires the Logres CLI: a root command with subcommands for
// build metadata and database migrations. The api/web/webAdmin surfaces are
// separate binaries (cmd/api, cmd/web, cmd/webAdmin) — this CLI is ops
// tooling only.
package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/wigata-intech/logres/internal/shared/cli"
)

// Execute builds the root command tree and dispatches os.Args. It returns an
// error suitable for mapping to a process exit code.
func Execute(ctx context.Context) error {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	root := cli.New("logres", "self-hostable, AGPL e-commerce platform")
	root.Sub(versionCommand())
	root.Sub(migrateCommand())

	return root.Run(ctx, os.Args[1:])
}

func versionCommand() *cli.Command {
	c := cli.New("version", "print build version")
	c.Exec = func(_ context.Context, _ []string) error {
		fmt.Printf("logres %s (commit %s, built %s)\n", cli.Version, cli.Commit, cli.Date)
		return nil
	}
	return c
}
