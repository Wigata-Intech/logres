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
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

const wantUpdateSQL = `UPDATE admin_users SET ` +
	`full_name = ?, email = ?, password = ?, status = ?, ` +
	`last_login_at = ?, updated_at = ?, deleted_at = ? ` +
	`WHERE id = ?`

func TestUpdate(t *testing.T) {
	t.Parallel()

	var (
		fixedID   = uuid.MustParse("018f0000-0000-7000-8000-000000000001")
		updated   = time.Date(2026, 1, 2, 3, 4, 6, 0, time.UTC)
		lastLogin = time.Date(2026, 1, 2, 3, 4, 7, 0, time.UTC)
		errBoom   = errors.New("exec failed")
		errRows   = errors.New("rows affected failed")
	)

	base := &model.AdminUser{
		ID: fixedID, FullName: adminFullName, Email: adminEmail,
		Password: adminPassword, Status: constant.AdminUserStatusActive,
		LastLoginAt: &lastLogin, UpdatedAt: updated,
	}

	type testCase struct {
		name    string
		input   *model.AdminUser
		setup   func(mock sqlmock.Sqlmock, in *model.AdminUser)
		wantErr error
	}

	cases := []testCase{
		{
			name:  "success updates one row",
			input: base,
			setup: func(mock sqlmock.Sqlmock, in *model.AdminUser) {
				mock.ExpectExec(wantUpdateSQL).WithArgs(
					in.FullName, in.Email, in.Password, in.Status,
					in.LastLoginAt, in.UpdatedAt, in.DeletedAt, in.ID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: nil,
		},
		{
			name:  "exec error is wrapped",
			input: base,
			setup: func(mock sqlmock.Sqlmock, _ *model.AdminUser) {
				mock.ExpectExec(wantUpdateSQL).WillReturnError(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name:  "rows-affected error is wrapped",
			input: base,
			setup: func(mock sqlmock.Sqlmock, _ *model.AdminUser) {
				mock.ExpectExec(wantUpdateSQL).
					WillReturnResult(sqlmock.NewErrorResult(errRows))
			},
			wantErr: errRows,
		},
		{
			name:  "no rows affected maps to ErrNotFound",
			input: base,
			setup: func(mock sqlmock.Sqlmock, _ *model.AdminUser) {
				mock.ExpectExec(wantUpdateSQL).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			wantErr: dbx.ErrNotFound,
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
			err = repo.Update(context.Background(), tc.input)

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
