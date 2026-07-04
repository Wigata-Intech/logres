package otp_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/shared/otp"
)

func TestGenerate(t *testing.T) {
	type testCase struct {
		name      string
		digits    int
		wantErr   bool
		wantMatch *regexp.Regexp
	}

	cases := []testCase{
		{name: "eight digits", digits: 8, wantMatch: regexp.MustCompile(`^\d{8}$`)},
		{name: "six digits", digits: 6, wantMatch: regexp.MustCompile(`^\d{6}$`)},
		{name: "single digit", digits: 1, wantMatch: regexp.MustCompile(`^\d$`)},
		{name: "zero is rejected", digits: 0, wantErr: true},
		{name: "negative is rejected", digits: -3, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := otp.Generate(tc.digits)
			if tc.wantErr {
				require.Error(t, err)
				assert.Empty(t, got)
				return
			}
			require.NoError(t, err)
			assert.Regexp(t, tc.wantMatch, got)
			assert.Len(t, got, tc.digits) // leading zeros preserved
		})
	}
}

func TestGenerateProducesVariety(t *testing.T) {
	seen := make(map[string]struct{})
	for range 50 {
		code, err := otp.Generate(8)
		require.NoError(t, err)
		seen[code] = struct{}{}
	}
	assert.Greater(t, len(seen), 40, "CSPRNG codes should rarely collide")
}

func TestHash(t *testing.T) {
	// Stable, hex-encoded SHA-256 of the code.
	assert.Equal(t,
		"ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f",
		otp.Hash("12345678"),
	)
	assert.Equal(t, otp.Hash("000000"), otp.Hash("000000"))
	assert.NotEqual(t, otp.Hash("000000"), otp.Hash("000001"))
	assert.Len(t, otp.Hash("anything"), 64)
}
