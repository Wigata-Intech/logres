package adminUser

import (
	"context"
	"fmt"
	"time"

	"github.com/wigata-intech/logres/internal/api/constant"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"
)

func (s *adminUserService) ChangePassword(ctx context.Context, req *sharedModel.AdminChangePasswordRequest) (*sharedModel.AdminChangePasswordResponse, error) {
	adminUser, err := s.adminUserRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.ChangePassword: %w", err)
	}
	if !adminUser.IsActive() {
		return nil, constant.ErrInvalidCredentials
	}

	ok, _, err := s.hasher.Verify(req.OldPassword, adminUser.Password)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.ChangePassword: verify: %w", err)
	}
	if !ok {
		return nil, constant.ErrInvalidCredentials
	}

	phc, err := s.hasher.Hash(req.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.ChangePassword: hash password: %w", err)
	}
	adminUser.Password = phc
	adminUser.UpdatedAt = time.Now()
	if err := s.adminUserRepo.Update(ctx, adminUser); err != nil {
		return nil, fmt.Errorf("adminUserService.ChangePassword: %w", err)
	}

	// TODO: [Tech Debt] revokes every session including the caller's own;
	// Phase 7's handler will pass the current refresh id so it can be excluded.
	if err := s.adminRefreshTokenRepo.RevokeAllForUser(ctx, adminUser.ID); err != nil {
		return nil, fmt.Errorf("adminUserService.ChangePassword: %w", err)
	}

	return &sharedModel.AdminChangePasswordResponse{}, nil
}
