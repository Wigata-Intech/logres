// Command logres-webadmin runs the Logres admin panel server.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/wigata-intech/logres/config"
	"github.com/wigata-intech/logres/internal/shared/cli"
	"github.com/wigata-intech/logres/internal/shared/envx"
	"github.com/wigata-intech/logres/internal/shared/httpx"
	"github.com/wigata-intech/logres/internal/webAdmin"
)

func main() {
	logger := slog.Default()
	if err := run(context.Background(), logger); err != nil {
		logger.Error("webadmin", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	if err := envx.LoadFile(".env", logger); err != nil {
		return fmt.Errorf("env: %w", err)
	}

	cfg, err := config.LoadWebAdmin()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	handler, err := webAdmin.New()
	if err != nil {
		return fmt.Errorf("build handler: %w", err)
	}

	logger.Info("logres-webadmin starting", "version", cli.Version)

	if err := httpx.Serve(ctx, ":"+cfg.Port, handler, logger); err != nil {
		return fmt.Errorf("serve: %w", err)
	}
	return nil
}
