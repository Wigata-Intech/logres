package cmd_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/cmd"
)

// runArgs invokes cmd.Execute with the given subcommand args, restoring
// os.Args afterward. Execute reads os.Args directly (it is the CLI entry
// point), so exercising subcommands means driving it the same way main does.
func runArgs(t *testing.T, args ...string) error {
	t.Helper()
	orig := os.Args
	t.Cleanup(func() { os.Args = orig })
	os.Args = append([]string{"logres"}, args...)
	return cmd.Execute(context.Background())
}

func TestExecuteVersion(t *testing.T) {
	err := runArgs(t, "version")
	require.NoError(t, err)
}

func TestExecuteMigrate(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "logres.db")
	migDir := filepath.Join(dir, "migration")

	t.Run("create", func(t *testing.T) {
		err := runArgs(t, "migrate", "create", "-dir", migDir, "widgets")
		require.NoError(t, err)

		entries, err := os.ReadDir(migDir)
		require.NoError(t, err)
		assert.Len(t, entries, 2)
	})

	t.Run("status before up", func(t *testing.T) {
		err := runArgs(t, "migrate", "status", "-db", dbPath, "-dir", migDir)
		require.NoError(t, err)
	})

	t.Run("up", func(t *testing.T) {
		err := runArgs(t, "migrate", "up", "-db", dbPath, "-dir", migDir)
		require.NoError(t, err)
	})

	t.Run("up again is a no-op", func(t *testing.T) {
		err := runArgs(t, "migrate", "up", "-db", dbPath, "-dir", migDir)
		require.NoError(t, err)
	})

	t.Run("down", func(t *testing.T) {
		err := runArgs(t, "migrate", "down", "-db", dbPath, "-dir", migDir)
		require.NoError(t, err)
	})

	t.Run("down with nothing left reports no-op", func(t *testing.T) {
		err := runArgs(t, "migrate", "down", "-db", dbPath, "-dir", migDir)
		require.NoError(t, err)
	})

	t.Run("status after down", func(t *testing.T) {
		err := runArgs(t, "migrate", "status", "-db", dbPath, "-dir", migDir)
		require.NoError(t, err)
	})
}

func TestExecuteMigrateCreateMissingName(t *testing.T) {
	dir := t.TempDir()
	err := runArgs(t, "migrate", "create", "-dir", filepath.Join(dir, "migration"))
	require.Error(t, err)
}

func TestExecuteUnknownCommand(t *testing.T) {
	err := runArgs(t, "does-not-exist")
	require.Error(t, err)
}
