package adminUser

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/wigata-intech/logres/internal/api/constant"
	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/shared/dbx"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"
)

func (s *adminUserService) Login(ctx context.Context, req *sharedModel.AdminLoginRequest) (*sharedModel.AdminTokenResponse, error) {
	adminUser, err := s.adminUserRepo.GetByEmail(ctx, req.Email)
	if errors.Is(err, dbx.ErrNotFound) {
		// Verify against a throwaway hash so an unknown email costs the same
		// argon2 work as a real one — anti-enumeration timing defense (OWASP
		// Authentication Cheat Sheet).
		if _, _, verifyErr := s.hasher.Verify(req.Password, dummyPHC); verifyErr != nil {
			return nil, fmt.Errorf("adminUserService.Login: dummy verify: %w", verifyErr)
		}
		return nil, constant.ErrInvalidCredentials
	}
	if err != nil {
		return nil, fmt.Errorf("adminUserService.Login: %w", err)
	}

	ok, needsRehash, err := s.hasher.Verify(req.Password, adminUser.Password)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.Login: verify: %w", err)
	}
	if !ok {
		return nil, constant.ErrInvalidCredentials
	}
	if !adminUser.IsActive() {
		return nil, constant.ErrInvalidCredentials
	}

	now := time.Now()
	if needsRehash {
		phc, err := s.hasher.Hash(req.Password)
		if err != nil {
			return nil, fmt.Errorf("adminUserService.Login: rehash: %w", err)
		}
		adminUser.Password = phc
	}
	adminUser.LastLoginAt = &now
	adminUser.UpdatedAt = now
	if err := s.adminUserRepo.Update(ctx, adminUser); err != nil {
		return nil, fmt.Errorf("adminUserService.Login: %w", err)
	}

	familyID, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("adminUserService.Login: family id: %w", err)
	}

	pair, err := s.issueTokenPair(ctx, adminUser.ID, familyID, model.AdminRefreshTokenDevice{})
	if err != nil {
		return nil, fmt.Errorf("adminUserService.Login: %w", err)
	}

	return &sharedModel.AdminTokenResponse{
		AccessToken:  pair.AccessToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(s.oauthTokenCfg.AccessTTL.Seconds()),
		RefreshToken: pair.RefreshToken,
	}, nil
}
