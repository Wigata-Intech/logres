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
	"github.com/wigata-intech/logres/internal/shared/token"

	mockrepository "github.com/wigata-intech/logres/mock/api/repository"
)

func TestRefreshToken(t *testing.T) {
	t.Parallel()

	userID := uuid.MustParse("018f0000-0000-7000-8000-000000000001")
	familyID := uuid.MustParse("018f0000-0000-7000-8000-000000000002")
	tokenID := uuid.MustParse("018f0000-0000-7000-8000-000000000003")

	req := &sharedModel.AdminRefreshTokenRequest{RefreshToken: "opaque-plaintext"}

	activeToken := func() *model.AdminRefreshToken {
		return &model.AdminRefreshToken{
			ID: tokenID, AdminUserID: userID, FamilyID: familyID,
			ExpiresAt: time.Now().Add(time.Hour),
		}
	}
	activeUser := func() *model.AdminUser {
		return &model.AdminUser{ID: userID, Status: constant.AdminUserStatusActive}
	}

	type mocks struct {
		userRepo    *mockrepository.MockIAdminUserRepository
		refreshRepo *mockrepository.MockIAdminRefreshTokenRepository
	}

	type testCase struct {
		name    string
		setup   func(m mocks)
		wantErr error
	}

	cases := []testCase{
		{
			name: "rotate ok",
			setup: func(m mocks) {
				rt := activeToken()
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(rt, nil)
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(activeUser(), nil)
				m.refreshRepo.EXPECT().Insert(mock.Anything, mock.MatchedBy(func(newRT *model.AdminRefreshToken) bool {
					return newRT.FamilyID == familyID
				})).Return(nil)
				m.refreshRepo.EXPECT().Rotate(mock.Anything, tokenID, mock.Anything).Return(nil)
			},
		},
		{
			name: "unknown token",
			setup: func(m mocks) {
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(nil, dbx.ErrNotFound)
			},
			wantErr: constant.ErrInvalidToken,
		},
		{
			name: "get by hash error",
			setup: func(m mocks) {
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(nil, errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "reused retired token revokes family",
			setup: func(m mocks) {
				rt := activeToken()
				revokedAt := time.Now().Add(-time.Minute)
				rt.RevokedAt = &revokedAt
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(rt, nil)
				m.refreshRepo.EXPECT().RevokeFamily(mock.Anything, familyID).Return(nil)
			},
			wantErr: constant.ErrInvalidToken,
		},
		{
			name: "revoke family error surfaces",
			setup: func(m mocks) {
				rt := activeToken()
				revokedAt := time.Now().Add(-time.Minute)
				rt.RevokedAt = &revokedAt
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(rt, nil)
				m.refreshRepo.EXPECT().RevokeFamily(mock.Anything, familyID).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "expired token",
			setup: func(m mocks) {
				rt := activeToken()
				rt.ExpiresAt = time.Now().Add(-time.Second)
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(rt, nil)
			},
			wantErr: constant.ErrInvalidToken,
		},
		{
			name: "deactivated user rejected",
			setup: func(m mocks) {
				rt := activeToken()
				u := activeUser()
				u.Status = constant.AdminUserStatusInactive
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(rt, nil)
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(u, nil)
			},
			wantErr: constant.ErrInvalidToken,
		},
		{
			name: "user not found rejected",
			setup: func(m mocks) {
				rt := activeToken()
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(rt, nil)
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(nil, dbx.ErrNotFound)
			},
			wantErr: constant.ErrInvalidToken,
		},
		{
			name: "get by id error surfaces",
			setup: func(m mocks) {
				rt := activeToken()
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(rt, nil)
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(nil, errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "rotate error",
			setup: func(m mocks) {
				rt := activeToken()
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(rt, nil)
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(activeUser(), nil)
				m.refreshRepo.EXPECT().Insert(mock.Anything, mock.Anything).Return(nil)
				m.refreshRepo.EXPECT().Rotate(mock.Anything, tokenID, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "insert error",
			setup: func(m mocks) {
				rt := activeToken()
				m.refreshRepo.EXPECT().GetByHash(mock.Anything, mock.Anything).Return(rt, nil)
				m.userRepo.EXPECT().GetByID(mock.Anything, userID).Return(activeUser(), nil)
				m.refreshRepo.EXPECT().Insert(mock.Anything, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			m := mocks{
				userRepo:    mockrepository.NewMockIAdminUserRepository(t),
				refreshRepo: mockrepository.NewMockIAdminRefreshTokenRepository(t),
			}
			tc.setup(m)

			issuer, err := token.New([]byte(testTokenSecret))
			require.NoError(t, err)

			svc := adminUser.New(
				adminUser.WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
				adminUser.WithAdminUserRepository(m.userRepo),
				adminUser.WithAdminRefreshTokenRepository(m.refreshRepo),
				adminUser.WithTokenIssuer(issuer),
				adminUser.WithOAuthTokenConfig(config.AdminOAuthTokenConfig{AccessTTL: time.Minute, RefreshTTL: time.Hour}),
			)

			got, err := svc.RefreshToken(context.Background(), req)

			if tc.wantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, got)
			assert.NotEmpty(t, got.AccessToken)
			assert.Equal(t, "Bearer", got.TokenType)
			assert.NotEmpty(t, got.RefreshToken)
		})
	}
}
