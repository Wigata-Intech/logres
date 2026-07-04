package adminUser_test

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/wigata-intech/logres/internal/shared/hasher"
	"github.com/wigata-intech/logres/internal/shared/token"
)

// testAdminEmail is the fixture email shared across the method test files —
// declared once so golangci-lint's goconst doesn't flag the repeat.
const testAdminEmail = "root@logres.dev"

// testTokenSecret is the fixed HMAC secret used by the real token issuer
// across the method test files — over 32 bytes to clear token.New's RFC 8725 floor.
const testTokenSecret = "test-secret-0123456789-abcdef-32b"

// errBoom is the generic repository/dependency failure used across the
// method test files to assert an error is wrapped and surfaced.
var errBoom = errors.New("boom")

// testHasher returns a real argon2id hasher with cost parameters low enough
// to keep the suite fast — these tests exercise behavior, not KDF cost.
func testHasher() hasher.Hasher {
	return hasher.New(hasher.WithParams(hasher.Params{Memory: 64, Time: 1, Threads: 1, SaltLen: 16, KeyLen: 32}))
}

// stubHasher is a hand-written test double (not mockery-generated) used only
// to force Hash/Verify error branches that a real argon2id hasher can't fail
// on demand.
type stubHasher struct {
	hashFn   func(string) (string, error)
	verifyFn func(string, string) (bool, bool, error)
}

func (s stubHasher) Hash(pw string) (string, error) { return s.hashFn(pw) }

func (s stubHasher) Verify(pw, enc string) (ok, needsRehash bool, err error) {
	return s.verifyFn(pw, enc)
}

// stubIssuer is a hand-written test double used only to force the token
// issuance error branch in issueTokenPair.
type stubIssuer struct{ err error }

func (s stubIssuer) Issue(uuid.UUID, time.Duration) (string, uuid.UUID, error) {
	return "", uuid.Nil, s.err
}

func (s stubIssuer) Verify(string) (*token.Claims, error) { return nil, s.err }
