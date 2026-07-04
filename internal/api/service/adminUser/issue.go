package adminUser

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/shared/otp"
)

// dummyPHC is a valid argon2id hash with no known plaintext. Login verifies an
// unknown email against it so the response takes the same argon2 cost whether
// or not the account exists — the standard anti-user-enumeration timing defense
// (OWASP Authentication Cheat Sheet; NIST SP 800-63B §5.2.2).
const dummyPHC = "$argon2id$v=19$m=65536,t=3,p=1$qKi2fJ0VMtJAPHXaz7sOTQ$jJ0Y6EOEXyjH5FrIFVfPOGCWvgyVCfnHnWpGGf/QPDw"

// refreshTokenBytes is the entropy of the opaque refresh-token plaintext:
// 256 bits, matching the access-token secret floor.
const refreshTokenBytes = 32

// hashToken is the at-rest representation of an opaque refresh-token
// plaintext. otp.Hash is reused because it is already the project's SHA-256
// hex helper — a refresh token has no separate hashing story.
func hashToken(plaintext string) string {
	return otp.Hash(plaintext)
}

func newOpaqueToken() (string, error) {
	buf := make([]byte, refreshTokenBytes)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("adminUserService.newOpaqueToken: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}

// tokenPair is the (access, refresh) pair issued on login and on rotation.
// RefreshTokenID lets a rotation caller link the old row to the new one
// without a redundant lookup.
type tokenPair struct {
	AccessToken    string
	RefreshToken   string
	RefreshTokenID uuid.UUID
}

// issueTokenPair mints an access JWT and a fresh opaque refresh token, then
// persists the refresh token under familyID. Callers pass an existing
// familyID to rotate within a login lineage, or a new one to start it.
func (s *adminUserService) issueTokenPair(
	ctx context.Context,
	userID, familyID uuid.UUID,
	device model.AdminRefreshTokenDevice,
) (*tokenPair, error) {
	accessToken, _, err := s.tokenIssuer.Issue(userID, s.oauthTokenCfg.AccessTTL)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.issueTokenPair: issue access token: %w", err)
	}

	refreshPlain, err := newOpaqueToken()
	if err != nil {
		return nil, fmt.Errorf("adminUserService.issueTokenPair: %w", err)
	}

	refreshToken, err := model.NewAdminRefreshToken(userID, hashToken(refreshPlain), familyID, s.oauthTokenCfg.RefreshTTL, device)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.issueTokenPair: new refresh token: %w", err)
	}

	if err := s.adminRefreshTokenRepo.Insert(ctx, refreshToken); err != nil {
		return nil, fmt.Errorf("adminUserService.issueTokenPair: insert refresh token: %w", err)
	}

	return &tokenPair{AccessToken: accessToken, RefreshToken: refreshPlain, RefreshTokenID: refreshToken.ID}, nil
}
