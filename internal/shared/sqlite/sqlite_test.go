package sqlite_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/shared/sqlite"
)

func TestOpen(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		cfg     sqlite.Config
		useDir  bool // point Path at a directory to force an open failure
		wantErr bool
	}

	cases := []testCase{
		{name: "empty config defaults to in-memory", cfg: sqlite.Config{}},
		{name: "explicit in-memory", cfg: sqlite.Config{Path: ":memory:", MaxOpenConns: 1}},
		{name: "shared-cache in-memory", cfg: sqlite.Config{Path: ":memory:", Shared: true}},
		{name: "unopenable path errors", useDir: true, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			cfg := tc.cfg
			if tc.useDir {
				cfg.Path = t.TempDir() // a directory cannot be opened as a database file
			}

			db, err := sqlite.Open(context.Background(), cfg)
			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, db)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, db)
			t.Cleanup(func() { _ = db.Close() })

			_, err = db.ExecContext(context.Background(), "CREATE TABLE t (id INTEGER PRIMARY KEY)")
			require.NoError(t, err)
			_, err = db.ExecContext(context.Background(), "INSERT INTO t (id) VALUES (1)")
			require.NoError(t, err)
		})
	}
}
