// Package cmd wires the Logres single-binary CLI: a root command with
// subcommands for build metadata and database migrations. Additional surfaces
// (api, web, webAdmin) attach here as they come online.
package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/wigata-intech/logres/internal/shared/cli"
)

// BuildInfo carries version metadata stamped into the binary at build time.
type BuildInfo struct {
	Version   string
	Commit    string
	BuildDate string
}

// Execute builds the root command tree and dispatches os.Args. It returns an
// error suitable for mapping to a process exit code.
func Execute(ctx context.Context, info BuildInfo) error {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	root := cli.New("logres", "self-hostable, AGPL e-commerce platform")
	root.Sub(versionCommand(info))
	root.Sub(migrateCommand())

	return root.Run(ctx, os.Args[1:])
}

func versionCommand(info BuildInfo) *cli.Command {
	c := cli.New("version", "print build version")
	c.Exec = func(_ context.Context, _ []string) error {
		fmt.Printf("logres %s (commit %s, built %s)\n", info.Version, info.Commit, info.BuildDate)
		return nil
	}
	return c
}
