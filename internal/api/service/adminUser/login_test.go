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
	"github.com/wigata-intech/logres/internal/shared/token"

	mockrepository "github.com/wigata-intech/logres/mock/api/repository"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	userID := uuid.MustParse("018f0000-0000-7000-8000-000000000001")

	req := &sharedModel.AdminLoginRequest{Email: testAdminEmail, Password: "s3cret"}

	h := testHasher()
	storedPHC, err := h.Hash(req.Password)
	require.NoError(t, err)
	wrongPHC, err := h.Hash("not-" + req.Password)
	require.NoError(t, err)
	// stalePHC hashes the right password under stronger params than the service's
	// current policy, so a successful verify also reports needsRehash.
	stalePHC, err := hasher.New(hasher.WithParams(hasher.Params{Memory: 64, Time: 2, Threads: 1, SaltLen: 16, KeyLen: 32})).Hash(req.Password)
	require.NoError(t, err)

	activeUser := func() *model.AdminUser {
		return &model.AdminUser{
			ID: userID, FullName: "Root Admin", Email: req.Email,
			Password: storedPHC, Status: constant.AdminUserStatusActive,
		}
	}

	type mocks struct {
		userRepo    *mockrepository.MockIAdminUserRepository
		refreshRepo *mockrepository.MockIAdminRefreshTokenRepository
		hasher      hasher.Hasher
		issuer      token.Issuer
	}

	type testCase struct {
		name    string
		setup   func(m *mocks)
		wantErr error
	}

	cases := []testCase{
		{
			name: "ok",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.MatchedBy(func(got *model.AdminUser) bool {
					return got.LastLoginAt != nil
				})).Return(nil)
				m.refreshRepo.EXPECT().Insert(mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "needs rehash rewrites stored hash",
			setup: func(m *mocks) {
				u := activeUser()
				u.Password = stalePHC
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.MatchedBy(func(got *model.AdminUser) bool {
					return got.Password != stalePHC && got.LastLoginAt != nil
				})).Return(nil)
				m.refreshRepo.EXPECT().Insert(mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "wrong password",
			setup: func(m *mocks) {
				u := activeUser()
				u.Password = wrongPHC
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
			},
			wantErr: constant.ErrInvalidCredentials,
		},
		{
			name: "unknown email uses dummy verify",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(nil, dbx.ErrNotFound)
			},
			wantErr: constant.ErrInvalidCredentials,
		},
		{
			name: "inactive user rejected",
			setup: func(m *mocks) {
				u := activeUser()
				u.Status = constant.AdminUserStatusInactive
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
			},
			wantErr: constant.ErrInvalidCredentials,
		},
		{
			name: "verify error",
			setup: func(m *mocks) {
				u := activeUser()
				u.Password = "not-a-valid-phc"
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
			},
			wantErr: hasher.ErrInvalidHash,
		},
		{
			name: "login get by email error",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(nil, errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "login update error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "refresh insert error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.Anything).Return(nil)
				m.refreshRepo.EXPECT().Insert(mock.Anything, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "dummy verify error",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(nil, dbx.ErrNotFound)
				m.hasher = stubHasher{verifyFn: func(string, string) (bool, bool, error) { return false, false, errBoom }}
			},
			wantErr: errBoom,
		},
		{
			name: "rehash error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.hasher = stubHasher{
					verifyFn: func(string, string) (bool, bool, error) { return true, true, nil },
					hashFn:   func(string) (string, error) { return "", errBoom },
				}
			},
			wantErr: errBoom,
		},
		{
			name: "issue token error",
			setup: func(m *mocks) {
				u := activeUser()
				m.userRepo.EXPECT().GetByEmail(mock.Anything, req.Email).Return(u, nil)
				m.userRepo.EXPECT().Update(mock.Anything, mock.Anything).Return(nil)
				m.issuer = stubIssuer{err: errBoom}
			},
			wantErr: errBoom,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			realIssuer, err := token.New([]byte(testTokenSecret))
			require.NoError(t, err)

			m := &mocks{
				userRepo:    mockrepository.NewMockIAdminUserRepository(t),
				refreshRepo: mockrepository.NewMockIAdminRefreshTokenRepository(t),
				hasher:      h,
				issuer:      realIssuer,
			}
			tc.setup(m)

			svc := adminUser.New(
				adminUser.WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
				adminUser.WithAdminUserRepository(m.userRepo),
				adminUser.WithAdminRefreshTokenRepository(m.refreshRepo),
				adminUser.WithHasher(m.hasher),
				adminUser.WithTokenIssuer(m.issuer),
				adminUser.WithOAuthTokenConfig(config.AdminOAuthTokenConfig{AccessTTL: time.Minute, RefreshTTL: time.Hour}),
			)

			got, err := svc.Login(context.Background(), req)

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
			assert.Equal(t, 60, got.ExpiresIn)
			assert.NotEmpty(t, got.RefreshToken)
		})
	}
}
