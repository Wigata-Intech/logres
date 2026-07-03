package adminUser_test

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/api/constant"
	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/api/repository/adminUser"
)

// wantInsertSQL is matched exactly (QueryMatcherEqual) so the test guards the
// statement's column list and order, not just the bound args.
const wantInsertSQL = `INSERT INTO admin_users ` +
	`(id, full_name, email, password, status, last_login_at, created_at, updated_at, deleted_at) ` +
	`VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

const (
	adminFullName = "Root Admin"
	adminPassword = "hashed"
	adminEmail    = "root@logres.dev"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	var (
		fixedID   = uuid.MustParse("018f0000-0000-7000-8000-000000000001")
		created   = time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
		updated   = time.Date(2026, 1, 2, 3, 4, 6, 0, time.UTC)
		lastLogin = time.Date(2026, 1, 2, 3, 4, 7, 0, time.UTC)
		errBoom   = errors.New("exec failed")
	)

	type testCase struct {
		name    string
		input   *model.AdminUser
		setup   func(mock sqlmock.Sqlmock, in *model.AdminUser)
		wantErr error
	}

	cases := []testCase{
		{
			name: "success with nullable timestamps present",
			input: &model.AdminUser{
				ID: fixedID, FullName: adminFullName, Email: adminEmail,
				Password: adminPassword, Status: constant.AdminUserStatusActive,
				LastLoginAt: &lastLogin, CreatedAt: created, UpdatedAt: updated,
			},
			setup: func(mock sqlmock.Sqlmock, in *model.AdminUser) {
				mock.ExpectExec(wantInsertSQL).WithArgs(
					in.ID, in.FullName, in.Email, in.Password, in.Status,
					in.LastLoginAt, in.CreatedAt, in.UpdatedAt, in.DeletedAt,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: nil,
		},
		{
			name: "success with nil timestamps",
			input: &model.AdminUser{
				ID: fixedID, FullName: "Root", Email: "root2@logres.dev",
				Password: adminPassword, Status: constant.AdminUserStatusInactive,
				CreatedAt: created, UpdatedAt: updated,
			},
			setup: func(mock sqlmock.Sqlmock, in *model.AdminUser) {
				mock.ExpectExec(wantInsertSQL).WithArgs(
					in.ID, in.FullName, in.Email, in.Password, in.Status,
					in.LastLoginAt, in.CreatedAt, in.UpdatedAt, in.DeletedAt,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: nil,
		},
		{
			name: "exec error is wrapped",
			input: &model.AdminUser{
				ID: fixedID, FullName: "Root", Email: adminEmail,
				Password: adminPassword, Status: constant.AdminUserStatusActive,
				CreatedAt: created, UpdatedAt: updated,
			},
			setup: func(mock sqlmock.Sqlmock, _ *model.AdminUser) {
				mock.ExpectExec(wantInsertSQL).WillReturnError(errBoom)
			},
			wantErr: errBoom,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			require.NoError(t, err)
			defer func() { _ = db.Close() }()

			tc.setup(mock, tc.input)

			repo := adminUser.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			err = repo.Create(context.Background(), tc.input)

			if tc.wantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.wantErr)
			} else {
				require.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
