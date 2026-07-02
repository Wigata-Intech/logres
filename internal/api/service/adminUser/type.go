package adminUser

import (
	"context"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
	"github.com/wigata-intech/logres/internal/api/service"
	"github.com/wigata-intech/logres/internal/shared/model"
)

type adminUserService struct {
	logger              *slog.Logger
	adminUserRepository repository.IAdminUserRepository
}

func New(
	logger *slog.Logger,
	adminUserRepository repository.IAdminUserRepository,
) service.IAdminUserService {
	return &adminUserService{
		logger:              logger,
		adminUserRepository: adminUserRepository,
	}
}

func (s *adminUserService) Setup(ctx context.Context, req *model.AdminSetupRequest) (*model.AdminSetupResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminUserService) Login(ctx context.Context, req *model.AdminLoginRequest) (*model.AdminTokenResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminUserService) RefreshToken(ctx context.Context, req *model.AdminRefreshTokenRequest) (*model.AdminTokenResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminUserService) ForgotPassword(ctx context.Context, req *model.AdminForgotPasswordRequest) (*model.AdminForgotPasswordResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminUserService) ChangePassword(ctx context.Context, req *model.AdminChangePasswordRequest) (*model.AdminChangePasswordResponse, error) {
	panic("not implemented") // TODO: Implement
}
