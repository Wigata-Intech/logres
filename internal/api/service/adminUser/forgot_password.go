package adminUser

import (
	"context"
	"errors"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/shared/dbx"
	"github.com/wigata-intech/logres/internal/shared/mailer"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"
	"github.com/wigata-intech/logres/internal/shared/otp"
)

// ForgotPassword issues a one-time reset code by email. It returns the same
// empty response whether or not the email exists (or is active), so a caller
// can't use this endpoint to enumerate accounts.
func (s *adminUserService) ForgotPassword(ctx context.Context, req *sharedModel.AdminForgotPasswordRequest) (*sharedModel.AdminForgotPasswordResponse, error) {
	resp := &sharedModel.AdminForgotPasswordResponse{}

	adminUser, err := s.adminUserRepo.GetByEmail(ctx, req.Email)
	if errors.Is(err, dbx.ErrNotFound) {
		return resp, nil
	}
	if err != nil {
		return nil, fmt.Errorf("adminUserService.ForgotPassword: %w", err)
	}
	if !adminUser.IsActive() {
		return resp, nil
	}

	code, err := otp.Generate(s.otpCfg.Length)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.ForgotPassword: generate otp: %w", err)
	}

	// Invalidate any prior unused code before creating the new one so
	// GetActiveByUser's LIMIT-1 stays correct: at most one active reset per
	// user+purpose.
	if err := s.adminResetPasswordRepo.InvalidateAllForUser(ctx, adminUser.ID, purposePasswordReset); err != nil {
		return nil, fmt.Errorf("adminUserService.ForgotPassword: %w", err)
	}

	reset, err := model.NewAdminPasswordReset(adminUser.ID, otp.Hash(code), purposePasswordReset, s.otpCfg.TTL)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.ForgotPassword: %w", err)
	}
	if err := s.adminResetPasswordRepo.Create(ctx, reset); err != nil {
		return nil, fmt.Errorf("adminUserService.ForgotPassword: %w", err)
	}

	msg := mailer.Message{
		To:      adminUser.Email,
		Subject: "Your password reset code",
		Text:    fmt.Sprintf("Your password reset code is %s. It expires in %s.", code, s.otpCfg.TTL),
	}
	if err := s.mailer.Send(ctx, msg); err != nil {
		return nil, fmt.Errorf("adminUserService.ForgotPassword: send mail: %w", err)
	}

	return resp, nil
}
