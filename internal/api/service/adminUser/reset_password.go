package adminUser

import (
	"context"
	"crypto/subtle"
	"errors"
	"fmt"
	"time"

	"github.com/wigata-intech/logres/internal/api/constant"
	"github.com/wigata-intech/logres/internal/shared/dbx"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"
	"github.com/wigata-intech/logres/internal/shared/otp"
)

// ResetPassword consumes an OTP issued by ForgotPassword. Every rejection
// path returns the same generic ErrInvalidReset — the caller can't tell an
// unknown email from an expired, used, locked, or wrong code.
func (s *adminUserService) ResetPassword(ctx context.Context, req *sharedModel.AdminResetPasswordRequest) (*sharedModel.AdminResetPasswordResponse, error) {
	adminUser, err := s.adminUserRepo.GetByEmail(ctx, req.Email)
	if errors.Is(err, dbx.ErrNotFound) {
		return nil, constant.ErrInvalidReset
	}
	if err != nil {
		return nil, fmt.Errorf("adminUserService.ResetPassword: %w", err)
	}

	reset, err := s.adminResetPasswordRepo.GetActiveByUser(ctx, adminUser.ID, purposePasswordReset)
	if errors.Is(err, dbx.ErrNotFound) {
		return nil, constant.ErrInvalidReset
	}
	if err != nil {
		return nil, fmt.Errorf("adminUserService.ResetPassword: %w", err)
	}

	// Expiry is checked here in Go, never in SQL: a time.Time in a DATETIME
	// column is stored with offset and does not order against datetime('now').
	if reset.UsedAt != nil || reset.AttemptCount >= s.otpCfg.AttemptMax || time.Now().After(reset.ExpiresAt) {
		return nil, constant.ErrInvalidReset
	}

	if subtle.ConstantTimeCompare([]byte(otp.Hash(req.OTP)), []byte(reset.OTPHash)) != 1 {
		if err := s.adminResetPasswordRepo.IncrementAttempt(ctx, reset.ID); err != nil {
			return nil, fmt.Errorf("adminUserService.ResetPassword: %w", err)
		}
		return nil, constant.ErrInvalidReset
	}

	phc, err := s.hasher.Hash(req.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("adminUserService.ResetPassword: hash password: %w", err)
	}
	adminUser.Password = phc
	adminUser.UpdatedAt = time.Now()
	if err := s.adminUserRepo.Update(ctx, adminUser); err != nil {
		return nil, fmt.Errorf("adminUserService.ResetPassword: %w", err)
	}

	if err := s.adminResetPasswordRepo.MarkUsed(ctx, reset.ID); err != nil {
		return nil, fmt.Errorf("adminUserService.ResetPassword: %w", err)
	}
	if err := s.adminRefreshTokenRepo.RevokeAllForUser(ctx, adminUser.ID); err != nil {
		return nil, fmt.Errorf("adminUserService.ResetPassword: %w", err)
	}

	return &sharedModel.AdminResetPasswordResponse{}, nil
}
