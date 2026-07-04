package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// minSecretLen is the RFC 8725 Section 3.5 floor for an HMAC key: at least as many bits
// as the hash output (256 bits for HS256).
const minSecretLen = 32

// ErrWeakSecret is returned by New when the signing secret is shorter than 256 bits.
var ErrWeakSecret = errors.New("token: signing secret must be at least 32 bytes")

// Claims are the access-token claims (RFC 9068). RegisteredClaims carries
// iss/sub/aud/exp/iat/jti; Subject and ID hold the admin id and jti.
type Claims struct {
	jwt.RegisteredClaims
}

// Issuer signs and verifies HS256 access tokens. Verify pins the algorithm to
// HS256, rejecting alg=none and RS256/HS256 confusion (RFC 8725).
type Issuer interface {
	Issue(subject uuid.UUID, ttl time.Duration) (token string, jti uuid.UUID, err error)
	Verify(token string) (*Claims, error)
}

type hs256Issuer struct {
	secret   []byte
	issuer   string
	audience string
	now      func() time.Time
}

func New(secret []byte, opts ...OptionFunc) (Issuer, error) {
	if len(secret) < minSecretLen {
		return nil, ErrWeakSecret
	}
	i := &hs256Issuer{secret: secret, now: time.Now}
	for _, opt := range opts {
		opt(i)
	}
	return i, nil
}

func (i *hs256Issuer) Issue(subject uuid.UUID, ttl time.Duration) (string, uuid.UUID, error) {
	jti, err := uuid.NewV7()
	if err != nil {
		return "", uuid.Nil, fmt.Errorf("token.Issue: jti: %w", err)
	}
	now := i.now()
	claims := Claims{jwt.RegisteredClaims{
		Issuer:    i.issuer,
		Subject:   subject.String(),
		Audience:  jwt.ClaimStrings{i.audience},
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        jti.String(),
	}}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tok.Header["typ"] = "at+jwt" // RFC 9068 Section 2.1

	signed, err := tok.SignedString(i.secret)
	if err != nil {
		return "", uuid.Nil, fmt.Errorf("token.Issue: sign: %w", err)
	}
	return signed, jti, nil
}

func (i *hs256Issuer) Verify(raw string) (*Claims, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(raw, &claims,
		func(*jwt.Token) (any, error) { return i.secret, nil },
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithIssuer(i.issuer),
		jwt.WithAudience(i.audience),
		jwt.WithExpirationRequired(),
		jwt.WithTimeFunc(i.now),
	)
	if err != nil {
		return nil, fmt.Errorf("token.Verify: %w", err)
	}
	return &claims, nil
}
