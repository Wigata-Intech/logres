package envx_test

import (
	"bytes"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/shared/envx"
)

func discardLogger() *slog.Logger { return slog.New(slog.NewTextHandler(io.Discard, nil)) }

const testKey = "ENVX_TEST_VALUE"

func TestGet(t *testing.T) {
	type testCase struct {
		name string
		val  string
		def  string
		want string
	}

	const def = "fallback"
	cases := []testCase{
		{name: "string override when set", val: "custom", def: def, want: "custom"},
		{name: "string falls back to default when unset", def: def, want: def},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(testKey, tc.val)

			got := envx.Get(testKey, tc.def)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDuration(t *testing.T) {
	type testCase struct {
		name    string
		val     string
		def     time.Duration
		want    time.Duration
		wantErr bool
	}

	cases := []testCase{
		{name: "duration override when set", val: "5m", def: time.Minute, want: 5 * time.Minute},
		{name: "duration falls back to default when unset", def: time.Minute, want: time.Minute},
		{name: "malformed duration returns error", val: "not-a-duration", def: time.Minute, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(testKey, tc.val)

			got, err := envx.Duration(testKey, tc.def)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestInt(t *testing.T) {
	type testCase struct {
		name    string
		val     string
		def     int
		want    int
		wantErr bool
	}

	cases := []testCase{
		{name: "int override when set", val: "42", def: 7, want: 42},
		{name: "int falls back to default when unset", def: 7, want: 7},
		{name: "malformed int returns error", val: "not-a-number", def: 7, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(testKey, tc.val)

			got, err := envx.Int(testKey, tc.def)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUint32(t *testing.T) {
	type testCase struct {
		name    string
		val     string
		def     uint32
		want    uint32
		wantErr bool
	}

	cases := []testCase{
		{name: "uint32 override when set", val: "65536", def: 1024, want: 65536},
		{name: "uint32 falls back to default when unset", def: 1024, want: 1024},
		{name: "malformed uint32 returns error", val: "not-a-number", def: 1024, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(testKey, tc.val)

			got, err := envx.Uint32(testKey, tc.def)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLoadFile(t *testing.T) {
	t.Run("loads unset keys, skips comments, real env wins", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), ".env")
		body := strings.Join([]string{
			"# a comment",
			"",
			"ENVX_LF_A=alpha",
			"export ENVX_LF_B=beta",
			`ENVX_LF_C="quoted value"`,
			"ENVX_LF_PRESET=fromfile",
		}, "\n")
		require.NoError(t, os.WriteFile(path, []byte(body), 0o600))

		t.Setenv("ENVX_LF_PRESET", "fromenv") // already in the env, must win
		for _, k := range []string{"ENVX_LF_A", "ENVX_LF_B", "ENVX_LF_C"} {
			t.Cleanup(func() { _ = os.Unsetenv(k) })
		}

		require.NoError(t, envx.LoadFile(path, discardLogger()))
		assert.Equal(t, "alpha", os.Getenv("ENVX_LF_A"))
		assert.Equal(t, "beta", os.Getenv("ENVX_LF_B"))
		assert.Equal(t, "quoted value", os.Getenv("ENVX_LF_C"))
		assert.Equal(t, "fromenv", os.Getenv("ENVX_LF_PRESET"))
	})

	t.Run("missing file is not an error", func(t *testing.T) {
		require.NoError(t, envx.LoadFile(filepath.Join(t.TempDir(), "absent.env"), discardLogger()))
	})

	t.Run("line without '=' errors without echoing content", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), ".env")
		require.NoError(t, os.WriteFile(path, []byte("SUPERSECRETTOKEN_NOEQ\n"), 0o600))
		err := envx.LoadFile(path, discardLogger())
		require.Error(t, err)
		assert.NotContains(t, err.Error(), "SUPERSECRETTOKEN", "parse error must not echo file content (may be a secret)")
	})

	t.Run("group/world-readable file warns", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("unix permission bits are meaningless on windows")
		}
		path := filepath.Join(t.TempDir(), ".env")
		require.NoError(t, os.WriteFile(path, []byte("ENVX_LF_PERM=ok\n"), 0o644))
		t.Cleanup(func() { _ = os.Unsetenv("ENVX_LF_PERM") })

		var buf bytes.Buffer
		logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelWarn}))
		require.NoError(t, envx.LoadFile(path, logger))
		assert.Contains(t, buf.String(), "group/world-readable")
	})
}
