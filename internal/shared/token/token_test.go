package token_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/shared/token"
)

const (
	testIssuer   = "logres-api"
	testAudience = "logres-admin"
)

var (
	testSecret = []byte("0123456789abcdef0123456789abcdef") // 32 bytes
	fixedNow   = time.Date(2026, 7, 3, 12, 0, 0, 0, time.UTC)
)

func newIssuer(t *testing.T, now time.Time) token.Issuer {
	t.Helper()
	iss, err := token.New(testSecret,
		token.WithIssuer(testIssuer),
		token.WithAudience(testAudience),
		token.WithClock(func() time.Time { return now }),
	)
	require.NoError(t, err)
	return iss
}

func TestNewRejectsWeakSecret(t *testing.T) {
	_, err := token.New([]byte("too-short"))
	assert.ErrorIs(t, err, token.ErrWeakSecret)
}

func TestIssueThenVerify(t *testing.T) {
	iss := newIssuer(t, fixedNow)
	sub := uuid.MustParse("018f0000-0000-7000-8000-000000000001")

	raw, jti, err := iss.Issue(sub, 15*time.Minute)
	require.NoError(t, err)
	assert.NotEmpty(t, raw)
	assert.NotEqual(t, uuid.Nil, jti)

	claims, err := iss.Verify(raw)
	require.NoError(t, err)
	assert.Equal(t, sub.String(), claims.Subject)
	assert.Equal(t, testIssuer, claims.Issuer)
	assert.Contains(t, claims.Audience, testAudience)
	assert.Equal(t, jti.String(), claims.ID)
	assert.True(t, claims.ExpiresAt.Equal(fixedNow.Add(15*time.Minute)))
	assert.True(t, claims.IssuedAt.Equal(fixedNow))
}

func TestIssueSetsAccessTokenType(t *testing.T) {
	iss := newIssuer(t, fixedNow)
	raw, _, err := iss.Issue(uuid.New(), time.Hour)
	require.NoError(t, err)

	parsed, _, err := jwt.NewParser().ParseUnverified(raw, jwt.MapClaims{})
	require.NoError(t, err)
	assert.Equal(t, "at+jwt", parsed.Header["typ"]) // RFC 9068
	assert.Equal(t, "HS256", parsed.Header["alg"])
}

func TestVerifyRejections(t *testing.T) {
	iss := newIssuer(t, fixedNow)
	sub := uuid.New()

	type testCase struct {
		name  string
		token func(t *testing.T) string
	}

	cases := []testCase{
		{
			name: "expired token",
			token: func(t *testing.T) string {
				raw, _, err := iss.Issue(sub, -time.Minute) // already expired at fixedNow
				require.NoError(t, err)
				return raw
			},
		},
		{
			name: "alg none",
			token: func(t *testing.T) string {
				tok := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.RegisteredClaims{
					Issuer:    testIssuer,
					Audience:  jwt.ClaimStrings{testAudience},
					ExpiresAt: jwt.NewNumericDate(fixedNow.Add(time.Hour)),
				})
				raw, err := tok.SignedString(jwt.UnsafeAllowNoneSignatureType)
				require.NoError(t, err)
				return raw
			},
		},
		{
			name: "wrong hmac algorithm",
			token: func(t *testing.T) string {
				tok := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
					Issuer:    testIssuer,
					Audience:  jwt.ClaimStrings{testAudience},
					ExpiresAt: jwt.NewNumericDate(fixedNow.Add(time.Hour)),
				})
				raw, err := tok.SignedString(testSecret)
				require.NoError(t, err)
				return raw
			},
		},
		{
			name: "tampered signature",
			token: func(t *testing.T) string {
				raw, _, err := iss.Issue(sub, time.Hour)
				require.NoError(t, err)
				return raw[:len(raw)-2] + "xx"
			},
		},
		{
			name: "wrong signing secret",
			token: func(t *testing.T) string {
				other, err := token.New([]byte("ffffffffffffffffffffffffffffffff"),
					token.WithIssuer(testIssuer),
					token.WithAudience(testAudience),
					token.WithClock(func() time.Time { return fixedNow }),
				)
				require.NoError(t, err)
				raw, _, err := other.Issue(sub, time.Hour)
				require.NoError(t, err)
				return raw
			},
		},
		{
			name: "wrong issuer",
			token: func(t *testing.T) string {
				other, err := token.New(testSecret,
					token.WithIssuer("someone-else"),
					token.WithAudience(testAudience),
					token.WithClock(func() time.Time { return fixedNow }),
				)
				require.NoError(t, err)
				raw, _, err := other.Issue(sub, time.Hour)
				require.NoError(t, err)
				return raw
			},
		},
		{
			name: "wrong audience",
			token: func(t *testing.T) string {
				other, err := token.New(testSecret,
					token.WithIssuer(testIssuer),
					token.WithAudience("other-aud"),
					token.WithClock(func() time.Time { return fixedNow }),
				)
				require.NoError(t, err)
				raw, _, err := other.Issue(sub, time.Hour)
				require.NoError(t, err)
				return raw
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			claims, err := iss.Verify(tc.token(t))
			require.Error(t, err)
			assert.Nil(t, claims)
		})
	}
}
