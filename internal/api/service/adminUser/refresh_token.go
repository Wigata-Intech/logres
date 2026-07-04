package adminUser

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wigata-intech/logres/internal/api/constant"
	"github.com/wigata-intech/logres/internal/shared/dbx"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"
)

// RefreshToken rotates an opaque refresh token for a new (access, refresh)
// pair. A retired token presented again is treated as reuse — the entire
// rotation family is revoked per RFC 9700 §4.14, forcing re-login.
func (s *adminUserService) RefreshToken(ctx context.Context, req *sharedModel.AdminRefreshTokenRequest) (*sharedModel.AdminTokenResponse, error) {
	hash := hashToken(req.RefreshToken)
	rt, err := s.adminRefreshTokenRepo.GetByHash(ctx, hash)
	if errors.Is(err, dbx.ErrNotFound) {
		return nil, constant.ErrInvalidToken
	}
	if err != nil {
		return nil, fmt.Errorf("adminUserService.RefreshToken: %w", err)
	}

	if rt.RevokedAt != nil {
		if err := s.adminRefreshTokenRepo.RevokeFamily(ctx, rt.FamilyID); err != nil {
			return nil, fmt.Errorf("adminUserService.RefreshToken: revoke reused family: %w", err)
		}
		return nil, constant.ErrInvalidToken
	}

	if time.Now().After(rt.ExpiresAt) {
		return nil, constant.ErrInvalidToken
	}

	// A valid token is worthless if the account was deactivated after it was
	// issued; re-check status so refresh can't outlive the login guard.
	adminUser, err := s.adminUserRepo.GetByID(ctx, rt.AdminUserID)
	if errors.Is(err, dbx.ErrNotFound) {
		return nil, constant.ErrInvalidToken
	}
	if err != nil {
		return nil, fmt.Errorf("adminUserService.RefreshToken: %w", err)
	}
	if !adminUser.IsActive() {
		return nil, constant.ErrInvalidToken
	}

	pair, err := s.issueTokenPair(ctx, rt.AdminUserID, rt.FamilyID, rt.Metadata.Device)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.RefreshToken: %w", err)
	}

	if err := s.adminRefreshTokenRepo.Rotate(ctx, rt.ID, pair.RefreshTokenID); err != nil {
		return nil, fmt.Errorf("adminUserService.RefreshToken: rotate: %w", err)
	}

	return &sharedModel.AdminTokenResponse{
		AccessToken:  pair.AccessToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(s.oauthTokenCfg.AccessTTL.Seconds()),
		RefreshToken: pair.RefreshToken,
	}, nil
}
