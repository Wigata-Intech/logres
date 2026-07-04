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
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"

	mockrepository "github.com/wigata-intech/logres/mock/api/repository"
	mockmailer "github.com/wigata-intech/logres/mock/shared/mailer"
)

func TestForgotPassword(t *testing.T) {
	t.Parallel()

	userID := uuid.MustParse("018f0000-0000-7000-8000-000000000001")

	req := &sharedModel.AdminForgotPasswordRequest{Email: testAdminEmail}

	activeUser := func() *model.AdminUser {
		return &model.AdminUser{ID: userID, Email: req.Email, Status: constant.AdminUserStatusActive}
	}

	type mocks struct {
		userRepo  *mockrepository.MockIAdminUserRepository
		resetRepo *mockrepository.MockIAdminPasswordResetRepository
		mailer    *mockmailer.MockMailer
		otpCfg    config.AdminOTPConfig
	}

	type testCase struct {
		name       string
		setup      func(m *mocks)
		wantErr    error
		wantErrMsg string
	}

	cases := []testCase{
		{
			name: "found invalidates before create then mails",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)

				var calls []string
				m.resetRepo.EXPECT().InvalidateAllForUser(mock.Anything, userID, mock.Anything).
					Run(func(context.Context, uuid.UUID, string) { calls = append(calls, "invalidate") }).
					Return(nil).Once()
				m.resetRepo.EXPECT().Create(mock.Anything, mock.Anything).
					Run(func(context.Context, *model.AdminPasswordReset) {
						calls = append(calls, "create")
						require.Equal(t, []string{"invalidate", "create"}, calls)
					}).Return(nil).Once()
				m.mailer.EXPECT().Send(mock.Anything, mock.MatchedBy(func(msg any) bool { return true })).Return(nil)
			},
		},
		{
			name: "unknown email sends no mail and creates nothing",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(nil, dbx.ErrNotFound)
			},
		},
		{
			name: "inactive user treated as unknown",
			setup: func(m *mocks) {
				u := activeUser()
				u.Status = constant.AdminUserStatusInactive
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
			},
		},
		{
			name: "forgot password get by email error",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(nil, errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "invalidate error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().InvalidateAllForUser(mock.Anything, userID, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "create error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().InvalidateAllForUser(mock.Anything, userID, mock.Anything).Return(nil)
				m.resetRepo.EXPECT().Create(mock.Anything, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "mail error surfaces",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.resetRepo.EXPECT().InvalidateAllForUser(mock.Anything, userID, mock.Anything).Return(nil)
				m.resetRepo.EXPECT().Create(mock.Anything, mock.Anything).Return(nil)
				m.mailer.EXPECT().Send(mock.Anything, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "otp generate error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.otpCfg = config.AdminOTPConfig{Length: 0}
			},
			wantErrMsg: "otp",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			m := &mocks{
				userRepo:  mockrepository.NewMockIAdminUserRepository(t),
				resetRepo: mockrepository.NewMockIAdminPasswordResetRepository(t),
				mailer:    mockmailer.NewMockMailer(t),
				otpCfg:    config.AdminOTPConfig{TTL: 10 * time.Minute, Length: 8},
			}
			tc.setup(m)

			svc := adminUser.New(
				adminUser.WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
				adminUser.WithAdminUserRepository(m.userRepo),
				adminUser.WithAdminPasswordResetRepository(m.resetRepo),
				adminUser.WithMailer(m.mailer),
				adminUser.WithOTPConfig(m.otpCfg),
			)

			got, err := svc.ForgotPassword(context.Background(), req)

			if tc.wantErrMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.wantErrMsg)
				assert.Nil(t, got)
				return
			}
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
