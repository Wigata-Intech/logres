package adminUser_test

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/api/constant"
	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/api/service/adminUser"
	"github.com/wigata-intech/logres/internal/shared/dbx"
	"github.com/wigata-intech/logres/internal/shared/hasher"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"

	mockrepository "github.com/wigata-intech/logres/mock/api/repository"
)

func TestChangePassword(t *testing.T) {
	t.Parallel()

	userID := uuid.MustParse("018f0000-0000-7000-8000-000000000001")

	req := &sharedModel.AdminChangePasswordRequest{
		UserID:      userID,
		OldPassword: "old-s3cret",
		NewPassword: "new-s3cret",
	}

	h := testHasher()
	storedPHC, err := h.Hash(req.OldPassword)
	require.NoError(t, err)
	wrongPHC, err := h.Hash("some-other-password")
	require.NoError(t, err)

	activeUser := func() *model.AdminUser {
		return &model.AdminUser{ID: userID, Password: storedPHC, Status: constant.AdminUserStatusActive}
	}

	type mocks struct {
		userRepo    *mockrepository.MockIAdminUserRepository
		refreshRepo *mockrepository.MockIAdminRefreshTokenRepository
		hasher      hasher.Hasher
	}

	type testCase struct {
		name    string
		setup   func(m *mocks)
		wantErr error
	}

	cases := []testCase{
		{
			name: "ok revokes other sessions",
			setup: func(m *mocks) {
				u := activeUser()
				originalPassword := u.Password
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(u, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.MatchedBy(func(got *model.AdminUser) bool {
					return got.Password != originalPassword
				})).Return(nil)
				m.refreshRepo.EXPECT().RevokeAllForUser(mock.Anything, userID).Return(nil)
			},
		},
		{
			name: "old password wrong",
			setup: func(m *mocks) {
				u := activeUser()
				u.Password = wrongPHC
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(u, nil)
			},
			wantErr: constant.ErrInvalidCredentials,
		},
		{
			name: "user not found",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(nil, dbx.ErrNotFound)
			},
			wantErr: dbx.ErrNotFound,
		},
		{
			name: "deactivated user rejected",
			setup: func(m *mocks) {
				u := activeUser()
				u.Status = constant.AdminUserStatusInactive
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(u, nil)
			},
			wantErr: constant.ErrInvalidCredentials,
		},
		{
			name: "verify error",
			setup: func(m *mocks) {
				u := activeUser()
				u.Password = "not-a-valid-phc"
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(u, nil)
			},
			wantErr: hasher.ErrInvalidHash,
		},
		{
			name: "change password update error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(u, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "revoke all error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(u, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.Anything).Return(nil)
				m.refreshRepo.EXPECT().RevokeAllForUser(mock.Anything, userID).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "hash new password error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(u, nil)
				m.hasher = stubHasher{
					verifyFn: func(string, string) (bool, bool, error) { return true, false, nil },
					hashFn:   func(string) (string, error) { return "", errBoom },
				}
			},
			wantErr: errBoom,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			m := &mocks{
				userRepo:    mockrepository.NewMockIAdminUserRepository(t),
				refreshRepo: mockrepository.NewMockIAdminRefreshTokenRepository(t),
				hasher:      h,
			}
			tc.setup(m)

			svc := adminUser.New(
				adminUser.WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
				adminUser.WithAdminUserRepository(m.userRepo),
				adminUser.WithAdminRefreshTokenRepository(m.refreshRepo),
				adminUser.WithHasher(m.hasher),
			)

			got, err := svc.ChangePassword(context.Background(), req)

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
