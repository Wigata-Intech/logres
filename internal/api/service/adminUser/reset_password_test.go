package adminUser_test

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/config"
	"github.com/wigata-intech/logres/internal/api/constant"
	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/api/service/adminUser"
	"github.com/wigata-intech/logres/internal/shared/dbx"
	"github.com/wigata-intech/logres/internal/shared/hasher"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"
	"github.com/wigata-intech/logres/internal/shared/otp"

	mockrepository "github.com/wigata-intech/logres/mock/api/repository"
)

func TestResetPassword(t *testing.T) {
	t.Parallel()

	userID := uuid.MustParse("018f0000-0000-7000-8000-000000000001")
	resetID := uuid.MustParse("018f0000-0000-7000-8000-000000000002")
	const validOTP = "12345678"

	req := &sharedModel.AdminResetPasswordRequest{
		Email:       testAdminEmail,
		OTP:         validOTP,
		NewPassword: "new-s3cret",
	}

	activeUser := func() *model.AdminUser {
		return &model.AdminUser{ID: userID, Email: req.Email, Password: "$argon2id$old"}
	}
	activeReset := func() *model.AdminPasswordReset {
		return &model.AdminPasswordReset{
			ID: resetID, AdminUserID: userID, OTPHash: otp.Hash(validOTP),
			AttemptCount: 0, ExpiresAt: time.Now().Add(time.Minute),
		}
	}

	type mocks struct {
		userRepo    *mockrepository.MockIAdminUserRepository
		resetRepo   *mockrepository.MockIAdminPasswordResetRepository
		refreshRepo *mockrepository.MockIAdminRefreshTokenRepository
		hasher      hasher.Hasher
	}

	type testCase struct {
		name    string
		req     *sharedModel.AdminResetPasswordRequest
		setup   func(m *mocks)
		wantErr error
	}

	cases := []testCase{
		{
			name: "valid otp updates password and revokes sessions",
			setup: func(m *mocks) {
				u := activeUser()
				originalPassword := u.Password
				r := activeReset()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.MatchedBy(func(got *model.AdminUser) bool {
					return got.Password != originalPassword
				})).Return(nil)
				m.resetRepo.EXPECT().MarkUsed(mock.Anything, resetID).Return(nil)
				m.refreshRepo.EXPECT().RevokeAllForUser(mock.Anything, userID).Return(nil)
			},
		},
		{
			name: "unknown email",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(nil, dbx.ErrNotFound)
			},
			wantErr: constant.ErrInvalidReset,
		},
		{
			name: "reset password get by email error",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(nil, errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "no active reset",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(nil, dbx.ErrNotFound)
			},
			wantErr: constant.ErrInvalidReset,
		},
		{
			name: "get active reset error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(nil, errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "already used",
			setup: func(m *mocks) {
				u := activeUser()
				r := activeReset()
				usedAt := time.Now().Add(-time.Second)
				r.UsedAt = &usedAt
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
			},
			wantErr: constant.ErrInvalidReset,
		},
		{
			name: "over attempt cap locked",
			setup: func(m *mocks) {
				u := activeUser()
				r := activeReset()
				r.AttemptCount = 5
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
			},
			wantErr: constant.ErrInvalidReset,
		},
		{
			name: "expired",
			setup: func(m *mocks) {
				u := activeUser()
				r := activeReset()
				r.ExpiresAt = time.Now().Add(-time.Second)
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
			},
			wantErr: constant.ErrInvalidReset,
		},
		{
			name: "wrong otp increments attempt",
			req:  &sharedModel.AdminResetPasswordRequest{Email: req.Email, OTP: "00000000", NewPassword: req.NewPassword},
			setup: func(m *mocks) {
				u := activeUser()
				r := activeReset()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
				m.resetRepo.EXPECT().IncrementAttempt(mock.Anything, resetID).Return(nil)
			},
			wantErr: constant.ErrInvalidReset,
		},
		{
			name: "increment attempt error",
			req:  &sharedModel.AdminResetPasswordRequest{Email: req.Email, OTP: "00000000", NewPassword: req.NewPassword},
			setup: func(m *mocks) {
				u := activeUser()
				r := activeReset()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
				m.resetRepo.EXPECT().IncrementAttempt(mock.Anything, resetID).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "reset password update error",
			setup: func(m *mocks) {
				u := activeUser()
				r := activeReset()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "mark used error",
			setup: func(m *mocks) {
				u := activeUser()
				r := activeReset()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.Anything).Return(nil)
				m.resetRepo.EXPECT().MarkUsed(mock.Anything, resetID).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "revoke all error",
			setup: func(m *mocks) {
				u := activeUser()
				r := activeReset()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.Anything).Return(nil)
				m.resetRepo.EXPECT().MarkUsed(mock.Anything, resetID).Return(nil)
				m.refreshRepo.EXPECT().RevokeAllForUser(mock.Anything, userID).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "hash new password error",
			setup: func(m *mocks) {
				u := activeUser()
				r := activeReset()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().GetActiveByUser(mock.Anything, userID, mock.Anything).Return(r, nil)
				m.hasher = stubHasher{hashFn: func(string) (string, error) { return "", errBoom }}
			},
			wantErr: errBoom,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			m := &mocks{
				userRepo:    mockrepository.NewMockIAdminUserRepository(t),
				resetRepo:   mockrepository.NewMockIAdminPasswordResetRepository(t),
				refreshRepo: mockrepository.NewMockIAdminRefreshTokenRepository(t),
				hasher:      testHasher(),
			}
			tc.setup(m)

			useReq := req
			if tc.req != nil {
				useReq = tc.req
			}

			svc := adminUser.New(
				adminUser.WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
				adminUser.WithAdminUserRepository(m.userRepo),
				adminUser.WithAdminPasswordResetRepository(m.resetRepo),
				adminUser.WithAdminRefreshTokenRepository(m.refreshRepo),
				adminUser.WithHasher(m.hasher),
				adminUser.WithOTPConfig(config.AdminOTPConfig{AttemptMax: 5}),
			)

			got, err := svc.ResetPassword(context.Background(), useReq)

			if tc.wantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, got)
		})
	}
}
