// Package envx parses environment variables into typed values with a
// default, so every config loader shares one implementation instead of
// re-typing os.Getenv + strconv per entity.
package envx

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Parse-error sentinels carry no line content — a value could be a secret, so
// callers report position (path:line), never the offending text.
var (
	errMissingEq = errors.New("missing '='")
	errEmptyKey  = errors.New("empty key")
)

const (
	// maxDotenvBytes caps the whole read; a config file is a few KB, so a larger
	// one is a mistake (wrong path) we refuse to slurp unbounded.
	maxDotenvBytes = 1 << 20 // 1 MiB
	// maxDotenvLine caps a single line so one monstrous line can't blow memory.
	maxDotenvLine = 1 << 16 // 64 KiB
)

// LoadFile reads a KEY=VALUE file (dotenv format) and sets each key that is not
// already present in the environment — so a real environment variable always
// wins over the file. A missing file is not an error (production sets env
// directly; the file is a local-dev convenience). Blank lines and #-comments
// are skipped; an optional leading "export " and surrounding quotes on the
// value are stripped. Because the file may hold secrets, it warns (via logger)
// when the file is group/world-readable and never echoes line content in parse
// errors — only the position (a value could be a secret). A nil logger falls
// back to slog.Default().
func LoadFile(path string, logger *slog.Logger) error {
	if logger == nil {
		logger = slog.Default()
	}

	f, err := os.Open(path) // #nosec G304 -- path is operator-controlled config, not user input
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("envx.LoadFile: %w", err)
	}
	defer func() { _ = f.Close() }()

	warnLoosePerms(f, path, logger)

	scanner := bufio.NewScanner(io.LimitReader(f, maxDotenvBytes))
	scanner.Buffer(make([]byte, 0, bufio.MaxScanTokenSize), maxDotenvLine)
	for line := 1; scanner.Scan(); line++ {
		if err := setDotenvLine(scanner.Text()); err != nil {
			return fmt.Errorf("envx.LoadFile: %s:%d: %w", path, line, err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("envx.LoadFile: %w", err)
	}
	return nil
}

// warnLoosePerms warns when a secrets-bearing file is readable by group/other
// (it should be 0600). It warns rather than refusing because deploy umasks
// vary; Unix perm bits are meaningless on Windows.
func warnLoosePerms(f *os.File, path string, logger *slog.Logger) {
	if runtime.GOOS == "windows" {
		return
	}
	info, err := f.Stat()
	if err != nil {
		return
	}
	if perm := info.Mode().Perm(); perm&0o077 != 0 {
		logger.Warn("dotenv file is group/world-readable; restrict it to 0600",
			"path", path, "perm", fmt.Sprintf("%04o", perm))
	}
}

// setDotenvLine parses one line and sets the key when it is not already in the
// environment. Blank and #-comment lines are no-ops. It returns a content-free
// sentinel on a malformed line so the caller can add position without leaking
// the (possibly secret) text.
func setDotenvLine(raw string) error {
	text := strings.TrimSpace(raw)
	if text == "" || strings.HasPrefix(text, "#") {
		return nil
	}
	text = strings.TrimPrefix(text, "export ")

	key, val, ok := strings.Cut(text, "=")
	if !ok {
		return errMissingEq
	}
	key = strings.TrimSpace(key)
	if key == "" {
		return errEmptyKey
	}
	val = strings.TrimSpace(val)
	if len(val) >= 2 && (val[0] == '"' || val[0] == '\'') && val[len(val)-1] == val[0] {
		val = val[1 : len(val)-1]
	}

	if _, set := os.LookupEnv(key); !set {
		return os.Setenv(key, val)
	}
	return nil
}

// Get returns the value of key, or def if unset or empty.
func Get(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// Duration parses key as a time.Duration, or returns def if unset.
func Duration(key string, def time.Duration) (time.Duration, error) {
	v := os.Getenv(key)
	if v == "" {
		return def, nil
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("%s: invalid duration %q: %w", key, v, err)
	}
	return d, nil
}

// Int parses key as an int, or returns def if unset.
func Int(key string, def int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return def, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("%s: invalid integer %q: %w", key, v, err)
	}
	return n, nil
}

// Uint32 parses key as a uint32, or returns def if unset.
func Uint32(key string, def uint32) (uint32, error) {
	v := os.Getenv(key)
	if v == "" {
		return def, nil
	}
	n, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("%s: invalid unsigned integer %q: %w", key, v, err)
	}
	return uint32(n), nil
}
