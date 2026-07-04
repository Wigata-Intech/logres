package adminUser

import (
	"log/slog"

	"github.com/wigata-intech/logres/config"
	"github.com/wigata-intech/logres/internal/api/repository"
	"github.com/wigata-intech/logres/internal/shared/hasher"
	"github.com/wigata-intech/logres/internal/shared/mailer"
	"github.com/wigata-intech/logres/internal/shared/token"
)

type OptionFunc func(*adminUserService)

func WithOAuthTokenConfig(cfg config.AdminOAuthTokenConfig) OptionFunc {
	return func(s *adminUserService) { s.oauthTokenCfg = cfg }
}

func WithOTPConfig(cfg config.AdminOTPConfig) OptionFunc {
	return func(s *adminUserService) { s.otpCfg = cfg }
}

func WithLogger(logger *slog.Logger) OptionFunc {
	return func(s *adminUserService) { s.logger = logger }
}

func WithHasher(h hasher.Hasher) OptionFunc {
	return func(s *adminUserService) { s.hasher = h }
}

func WithTokenIssuer(issuer token.Issuer) OptionFunc {
	return func(s *adminUserService) { s.tokenIssuer = issuer }
}

func WithMailer(m mailer.Mailer) OptionFunc {
	return func(s *adminUserService) { s.mailer = m }
}

func WithAdminUserRepository(repo repository.IAdminUserRepository) OptionFunc {
	return func(s *adminUserService) { s.adminUserRepo = repo }
}

func WithAdminPasswordResetRepository(repo repository.IAdminPasswordResetRepository) OptionFunc {
	return func(s *adminUserService) { s.adminResetPasswordRepo = repo }
}

func WithAdminRefreshTokenRepository(repo repository.IAdminRefreshTokenRepository) OptionFunc {
	return func(s *adminUserService) { s.adminRefreshTokenRepo = repo }
}
