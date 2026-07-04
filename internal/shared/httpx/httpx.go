// Package httpx holds the graceful-serve helper shared by every Logres HTTP
// binary (api, web, webAdmin).
package httpx

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const shutdownTimeout = 10 * time.Second

// Serve runs handler on addr until ctx is cancelled or an interrupt/TERM
// signal arrives, then shuts down gracefully. It returns nil on a clean
// shutdown, or the first error either ListenAndServe or Shutdown produced.
func Serve(ctx context.Context, addr string, handler http.Handler, logger *slog.Logger) error {
	srv := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	serveErr := make(chan error, 1)
	go func() {
		logger.Info("server starting", "addr", srv.Addr)
		serveErr <- srv.ListenAndServe()
	}()

	select {
	case err := <-serveErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("httpx.Serve: %w", err)
		}
		return nil
	case <-ctx.Done():
		logger.Info("server shutting down", "addr", srv.Addr)
		// ctx is already Done (that's why we're here); WithoutCancel keeps it
		// as the parent without inheriting its expired deadline.
		shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), shutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("httpx.Serve: shutdown: %w", err)
		}
		logger.Info("server stopped", "addr", srv.Addr)
		return nil
	}
}
