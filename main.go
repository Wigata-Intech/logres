package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/wigata-intech/logres/cmd"
	"github.com/wigata-intech/logres/internal/shared/cli"
)

// Build metadata, overridden via -ldflags at build time (see Makefile).
var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
)

func main() {
	err := cmd.Execute(context.Background(), cmd.BuildInfo{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,
	})
	if err == nil {
		return
	}

	if errors.Is(err, cli.ErrUsage) {
		os.Exit(2)
	}
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
