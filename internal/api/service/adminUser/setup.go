package adminUser

import (
	"context"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/constant"
	"github.com/wigata-intech/logres/internal/api/model"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"
)

// Setup provisions the single admin account a fresh instance boots with. It
// is only callable once — a second call fails with ErrAdminAlreadyExists so
// an exposed setup endpoint can't be used to mint extra admins.
func (s *adminUserService) Setup(ctx context.Context, req *sharedModel.AdminSetupRequest) (*sharedModel.AdminSetupResponse, error) {
	count, err := s.adminUserRepo.CountAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.Setup: %w", err)
	}
	if count > 0 {
		return nil, constant.ErrAdminAlreadyExists
	}

	phc, err := s.hasher.Hash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.Setup: hash password: %w", err)
	}

	adminUser, err := model.NewAdminUser(req.FullName, req.Email)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.Setup: %w", err)
	}
	adminUser.Password = phc
	adminUser.Status = constant.AdminUserStatusActive

	if err := s.adminUserRepo.Create(ctx, adminUser); err != nil {
		return nil, fmt.Errorf("adminUserService.Setup: %w", err)
	}

	return &sharedModel.AdminSetupResponse{}, nil
}
