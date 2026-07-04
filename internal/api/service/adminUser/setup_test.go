package adminUser_test

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	mockrepository "github.com/wigata-intech/logres/mock/api/repository"

	"github.com/wigata-intech/logres/internal/api/constant"
	"github.com/wigata-intech/logres/internal/api/service/adminUser"
	"github.com/wigata-intech/logres/internal/shared/hasher"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"
)

func TestSetup(t *testing.T) {
	t.Parallel()

	req := &sharedModel.AdminSetupRequest{
		FullName: "Root Admin",
		Email:    testAdminEmail,
		Password: "correct horse battery staple",
	}

	type mocks struct {
		userRepo *mockrepository.MockIAdminUserRepository
		hasher   hasher.Hasher
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
				m.userRepo.EXPECT().CountAll(mock.Anything).Return(0, nil)
				m.userRepo.EXPECT().Create(mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "already exists",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().CountAll(mock.Anything).Return(1, nil)
			},
			wantErr: constant.ErrAdminAlreadyExists,
		},
		{
			name: "count error",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().CountAll(mock.Anything).Return(0, errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "create error",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().CountAll(mock.Anything).Return(0, nil)
				m.userRepo.EXPECT().Create(mock.Anything, mock.Anything).Return(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "hash error",
			setup: func(m *mocks) {
				m.userRepo.EXPECT().CountAll(mock.Anything).Return(0, nil)
				m.hasher = stubHasher{hashFn: func(string) (string, error) { return "", errBoom }}
			},
			wantErr: errBoom,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			m := &mocks{
				userRepo: mockrepository.NewMockIAdminUserRepository(t),
				hasher:   testHasher(),
			}
			tc.setup(m)

			svc := adminUser.New(
				adminUser.WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
				adminUser.WithAdminUserRepository(m.userRepo),
				adminUser.WithHasher(m.hasher),
			)

			got, err := svc.Setup(context.Background(), req)

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
