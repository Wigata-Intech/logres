package adminPasswordReset_test

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

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/api/repository/adminPasswordReset"
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

const wantSelectActiveSQL = `SELECT ` +
	`id, admin_user_id, otp_hash, purpose, attempt_count, expires_at, used_at, created_at ` +
	`FROM admin_password_resets ` +
	`WHERE admin_user_id = ? AND purpose = ? AND used_at IS NULL ` +
	`ORDER BY created_at DESC LIMIT 1`

var resetColumns = []string{
	"id", "admin_user_id", "otp_hash", "purpose",
	"attempt_count", "expires_at", "used_at", "created_at",
}

func TestGetActiveByUser(t *testing.T) {
	t.Parallel()

	const (
		idText          = "018f0000-0000-7000-8000-000000000001"
		adminUserIDText = "018f0000-0000-7000-8000-000000000002"
	)
	var (
		fixedID     = uuid.MustParse(idText)
		adminUserID = uuid.MustParse(adminUserIDText)
		expires     = time.Date(2026, 1, 2, 3, 14, 5, 0, time.UTC)
		created     = time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
		errBoom     = errors.New("query failed")
	)

	type testCase struct {
		name       string
		setup      func(mock sqlmock.Sqlmock)
		wantErrIs  error
		wantResult *model.AdminPasswordReset
	}

	cases := []testCase{
		{
			name: "found active reset",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows(resetColumns).AddRow(
					idText, adminUserIDText, resetOTPHash, resetPurpose,
					int64(0), expires, nil, created,
				)
				mock.ExpectQuery(wantSelectActiveSQL).WithArgs(adminUserID, resetPurpose).WillReturnRows(rows)
			},
			wantResult: &model.AdminPasswordReset{
				ID: fixedID, AdminUserID: adminUserID, OTPHash: resetOTPHash,
				Purpose: resetPurpose, ExpiresAt: expires, CreatedAt: created,
			},
		},
		{
			name: "not found maps to ErrNotFound",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(wantSelectActiveSQL).WithArgs(adminUserID, resetPurpose).WillReturnError(sql.ErrNoRows)
			},
			wantErrIs: dbx.ErrNotFound,
		},
		{
			name: "scan error is wrapped",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(wantSelectActiveSQL).WithArgs(adminUserID, resetPurpose).WillReturnError(errBoom)
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

			repo := adminPasswordReset.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			got, err := repo.GetActiveByUser(context.Background(), adminUserID, resetPurpose)

			if tc.wantErrIs != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.wantErrIs)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
				assertPasswordResetEqual(t, tc.wantResult, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func assertPasswordResetEqual(t *testing.T, want, got *model.AdminPasswordReset) {
	t.Helper()
	assert.Equal(t, want.ID, got.ID)
	assert.Equal(t, want.AdminUserID, got.AdminUserID)
	assert.Equal(t, want.OTPHash, got.OTPHash)
	assert.Equal(t, want.Purpose, got.Purpose)
	assert.Equal(t, want.AttemptCount, got.AttemptCount)
	assert.True(t, want.ExpiresAt.Equal(got.ExpiresAt))
	assert.True(t, want.CreatedAt.Equal(got.CreatedAt))
	if want.UsedAt == nil {
		assert.Nil(t, got.UsedAt)
	} else {
		require.NotNil(t, got.UsedAt)
		assert.True(t, want.UsedAt.Equal(*got.UsedAt))
	}
}
