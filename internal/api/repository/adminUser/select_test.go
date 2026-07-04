package adminUser_test

import (
	"context"
	"database/sql"
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

const wantSelectByEmailSQL = `SELECT ` +
	`id, full_name, email, password, status, last_login_at, created_at, updated_at, deleted_at ` +
	`FROM admin_users WHERE email = ? AND deleted_at IS NULL`

const wantSelectByIDSQL = `SELECT ` +
	`id, full_name, email, password, status, last_login_at, created_at, updated_at, deleted_at ` +
	`FROM admin_users WHERE id = ? AND deleted_at IS NULL`

var adminUserColumns = []string{
	"id", "full_name", "email", "password", "status",
	"last_login_at", "created_at", "updated_at", "deleted_at",
}

func TestGetByEmail(t *testing.T) {
	t.Parallel()

	const idText = "018f0000-0000-7000-8000-000000000001"
	var (
		fixedID   = uuid.MustParse(idText)
		created   = time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
		updated   = time.Date(2026, 1, 2, 3, 4, 6, 0, time.UTC)
		lastLogin = time.Date(2026, 1, 2, 3, 4, 7, 0, time.UTC)
		errBoom   = errors.New("query failed")
	)

	type testCase struct {
		name       string
		setup      func(mock sqlmock.Sqlmock)
		wantErrIs  error
		wantResult *model.AdminUser
	}

	cases := []testCase{
		{
			name: "found active user",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows(adminUserColumns).AddRow(
					idText, adminFullName, adminEmail, adminPassword, int64(1),
					lastLogin, created, updated, nil,
				)
				mock.ExpectQuery(wantSelectByEmailSQL).WithArgs(adminEmail).WillReturnRows(rows)
			},
			wantResult: &model.AdminUser{
				ID: fixedID, FullName: adminFullName, Email: adminEmail,
				Password: adminPassword, Status: constant.AdminUserStatusActive,
				LastLoginAt: &lastLogin, CreatedAt: created, UpdatedAt: updated,
			},
		},
		{
			name: "not found maps to ErrNotFound",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(wantSelectByEmailSQL).WithArgs(adminEmail).WillReturnError(sql.ErrNoRows)
			},
			wantErrIs: dbx.ErrNotFound,
		},
		{
			name: "scan error is wrapped",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(wantSelectByEmailSQL).WithArgs(adminEmail).WillReturnError(errBoom)
			},
			wantErrIs: errBoom,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			require.NoError(t, err)
			defer func() { _ = db.Close() }()

			tc.setup(mock)

			repo := adminUser.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			got, err := repo.GetByEmail(context.Background(), adminEmail)

			if tc.wantErrIs != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.wantErrIs)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
				assertAdminUserEqual(t, tc.wantResult, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetByID(t *testing.T) {
	t.Parallel()

	const idText = "018f0000-0000-7000-8000-000000000001"
	var (
		fixedID   = uuid.MustParse(idText)
		created   = time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
		updated   = time.Date(2026, 1, 2, 3, 4, 6, 0, time.UTC)
		lastLogin = time.Date(2026, 1, 2, 3, 4, 7, 0, time.UTC)
		errBoom   = errors.New("query failed")
	)

	type testCase struct {
		name       string
		setup      func(mock sqlmock.Sqlmock)
		wantErrIs  error
		wantResult *model.AdminUser
	}

	cases := []testCase{
		{
			name: "found active user",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows(adminUserColumns).AddRow(
					idText, adminFullName, adminEmail, adminPassword, int64(1),
					lastLogin, created, updated, nil,
				)
				mock.ExpectQuery(wantSelectByIDSQL).WithArgs(fixedID).WillReturnRows(rows)
			},
			wantResult: &model.AdminUser{
				ID: fixedID, FullName: adminFullName, Email: adminEmail,
				Password: adminPassword, Status: constant.AdminUserStatusActive,
				LastLoginAt: &lastLogin, CreatedAt: created, UpdatedAt: updated,
			},
		},
		{
			name: "not found maps to ErrNotFound",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(wantSelectByIDSQL).WithArgs(fixedID).WillReturnError(sql.ErrNoRows)
			},
			wantErrIs: dbx.ErrNotFound,
		},
		{
			name: "scan error is wrapped",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(wantSelectByIDSQL).WithArgs(fixedID).WillReturnError(errBoom)
			},
			wantErrIs: errBoom,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			require.NoError(t, err)
			defer func() { _ = db.Close() }()

			tc.setup(mock)

			repo := adminUser.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			got, err := repo.GetByID(context.Background(), fixedID)

			if tc.wantErrIs != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.wantErrIs)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
				assertAdminUserEqual(t, tc.wantResult, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// assertAdminUserEqual compares timestamps with time.Equal so wall-clock instants
// match regardless of internal encoding.
func assertAdminUserEqual(t *testing.T, want, got *model.AdminUser) {
	t.Helper()
	assert.Equal(t, want.ID, got.ID)
	assert.Equal(t, want.FullName, got.FullName)
	assert.Equal(t, want.Email, got.Email)
	assert.Equal(t, want.Password, got.Password)
	assert.Equal(t, want.Status, got.Status)
	assert.True(t, want.CreatedAt.Equal(got.CreatedAt))
	assert.True(t, want.UpdatedAt.Equal(got.UpdatedAt))
	assertTimePtrEqual(t, want.LastLoginAt, got.LastLoginAt)
	assertTimePtrEqual(t, want.DeletedAt, got.DeletedAt)
}

func assertTimePtrEqual(t *testing.T, want, got *time.Time) {
	t.Helper()
	if want == nil {
		assert.Nil(t, got)
		return
	}
	require.NotNil(t, got)
	assert.True(t, want.Equal(*got))
}
