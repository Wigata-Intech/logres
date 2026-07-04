package adminUser

import (
	"log/slog"

	"github.com/wigata-intech/logres/config"
	"github.com/wigata-intech/logres/internal/api/repository"
	"github.com/wigata-intech/logres/internal/api/service"
	"github.com/wigata-intech/logres/internal/shared/hasher"
	"github.com/wigata-intech/logres/internal/shared/mailer"
	"github.com/wigata-intech/logres/internal/shared/token"
)

const purposePasswordReset = "password_reset"

type adminUserService struct {
	oauthTokenCfg          config.AdminOAuthTokenConfig
	otpCfg                 config.AdminOTPConfig
	logger                 *slog.Logger
	hasher                 hasher.Hasher
	tokenIssuer            token.Issuer
	mailer                 mailer.Mailer
	adminUserRepo          repository.IAdminUserRepository
	adminResetPasswordRepo repository.IAdminPasswordResetRepository
	adminRefreshTokenRepo  repository.IAdminRefreshTokenRepository
}

func New(opts ...OptionFunc) service.IAdminUserService {
	s := &adminUserService{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
