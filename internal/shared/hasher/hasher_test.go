package hasher_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/shared/hasher"
)

// testParams keep argon2 cheap so the suite stays fast; production uses DefaultParams.
var testParams = hasher.Params{Memory: 8 * 1024, Time: 1, Threads: 1, SaltLen: 16, KeyLen: 32}

var phcPattern = regexp.MustCompile(`^\$argon2id\$v=19\$m=\d+,t=\d+,p=\d+\$[^$]+\$[^$]+$`)

func TestHashProducesPHCString(t *testing.T) {
	h := hasher.New(hasher.WithParams(testParams))

	encoded, err := h.Hash("correct horse battery staple")
	require.NoError(t, err)
	assert.Regexp(t, phcPattern, encoded)

	// Distinct salts make every hash of the same password unique.
	other, err := h.Hash("correct horse battery staple")
	require.NoError(t, err)
	assert.NotEqual(t, encoded, other)
}

func TestVerify(t *testing.T) {
	h := hasher.New(hasher.WithParams(testParams))
	encoded, err := h.Hash("s3cret-password")
	require.NoError(t, err)

	type testCase struct {
		name            string
		password        string
		encoded         string
		wantOK          bool
		wantNeedsRehash bool
		wantErrIs       error
	}

	cases := []testCase{
		{name: "correct password", password: "s3cret-password", encoded: encoded, wantOK: true},
		{name: "wrong password", password: "not-it", encoded: encoded, wantOK: false},
		{name: "malformed encoding", password: "x", encoded: "not-a-phc-string", wantErrIs: hasher.ErrInvalidHash},
		{name: "wrong field count", password: "x", encoded: "$argon2id$v=19$m=8,t=1$salt", wantErrIs: hasher.ErrInvalidHash},
		{name: "wrong variant", password: "x", encoded: "$argon2i$v=19$m=8,t=1,p=1$c2FsdA$aGFzaA", wantErrIs: hasher.ErrInvalidHash},
		{name: "trailing garbage in params", password: "x", encoded: "$argon2id$v=19$m=8,t=1,p=1;DROP$c2FsdA$aGFzaA", wantErrIs: hasher.ErrInvalidHash},
		{name: "bad base64 salt", password: "x", encoded: "$argon2id$v=19$m=8,t=1,p=1$!!!$aGFzaA", wantErrIs: hasher.ErrInvalidHash},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ok, needsRehash, err := h.Verify(tc.password, tc.encoded)
			if tc.wantErrIs != nil {
				assert.ErrorIs(t, err, tc.wantErrIs)
				assert.False(t, ok)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantNeedsRehash, needsRehash)
		})
	}
}

func TestVerifyReportsNeedsRehashWhenParamsChange(t *testing.T) {
	// Hash under the low params, then verify with a hasher on stronger params.
	old := hasher.New(hasher.WithParams(testParams))
	encoded, err := old.Hash("rolling-forward")
	require.NoError(t, err)

	stronger := hasher.Params{Memory: 16 * 1024, Time: 2, Threads: 1, SaltLen: 16, KeyLen: 32}
	current := hasher.New(hasher.WithParams(stronger))

	ok, needsRehash, err := current.Verify("rolling-forward", encoded)
	require.NoError(t, err)
	assert.True(t, ok)
	assert.True(t, needsRehash)

	// Re-hashing under the current policy clears the flag.
	rehashed, err := current.Hash("rolling-forward")
	require.NoError(t, err)
	ok, needsRehash, err = current.Verify("rolling-forward", rehashed)
	require.NoError(t, err)
	assert.True(t, ok)
	assert.False(t, needsRehash)
}
