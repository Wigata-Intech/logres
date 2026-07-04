package main

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/config"
	"github.com/wigata-intech/logres/internal/shared/sqlite"
)

func TestRun(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	t.Run("fails fast on missing jwt secret", func(t *testing.T) {
		t.Setenv("AUTH_JWT_SECRET", "")

		err := run(context.Background(), logger)
		require.Error(t, err)
	})

	t.Run("fails on unopenable db path", func(t *testing.T) {
		t.Setenv("AUTH_JWT_SECRET", "01234567890123456789012345678901")
		t.Setenv("DB_PATH", t.TempDir()) // a directory is never a valid sqlite file

		err := run(context.Background(), logger)
		require.Error(t, err)
	})

	t.Run("serves until shutdown", func(t *testing.T) {
		t.Setenv("AUTH_JWT_SECRET", "01234567890123456789012345678901")
		t.Setenv("DB_PATH", ":memory:")
		t.Setenv("PORT", "0") // ephemeral: any free port, no collision with other tests

		// A context that's already gone by the time the server would block on
		// ListenAndServe drives the same ctx.Done() shutdown path a real
		// SIGTERM would, without needing an OS signal or a fixed port.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		err := run(ctx, logger)
		require.NoError(t, err)
	})
}

func TestBuildAPIHandler(t *testing.T) {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	db, err := sqlite.Open(ctx, sqlite.Config{Path: ":memory:", MaxOpenConns: 1})
	require.NoError(t, err)
	defer func() { _ = db.Close() }()

	cfg := &config.Config{
		API: config.APIServerConfig{Port: "8080"},
		DB:  config.DBConfig{Path: ":memory:"},
		Auth: config.AuthConfig{
			JWTSecret: "01234567890123456789012345678901",
			Argon2: config.Argon2Config{
				Memory:      65536,
				Time:        3,
				Threads:     1,
				Concurrency: 4,
			},
		},
	}

	handler, err := buildAPIHandler(logger, db, cfg)
	require.NoError(t, err)
	assert.NotNil(t, handler)
}
