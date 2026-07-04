package adminPasswordReset_test

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

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/api/repository/adminPasswordReset"
)

const caseExecErr = "exec error is wrapped"

// wantInsertSQL is matched exactly (QueryMatcherEqual) so the test guards the
// statement's column list and order, not just the bound args.
const wantInsertSQL = `INSERT INTO admin_password_resets ` +
	`(id, admin_user_id, otp_hash, purpose, attempt_count, expires_at, used_at, created_at) ` +
	`VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

const (
	resetOTPHash = "deadbeef"
	resetPurpose = "password_reset"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	var (
		fixedID     = uuid.MustParse("018f0000-0000-7000-8000-000000000001")
		adminUserID = uuid.MustParse("018f0000-0000-7000-8000-000000000002")
		expires     = time.Date(2026, 1, 2, 3, 14, 5, 0, time.UTC)
		created     = time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
		usedAt      = time.Date(2026, 1, 2, 3, 5, 0, 0, time.UTC)
		errBoom     = errors.New("exec failed")
	)

	type testCase struct {
		name    string
		input   *model.AdminPasswordReset
		setup   func(mock sqlmock.Sqlmock, in *model.AdminPasswordReset)
		wantErr error
	}

	cases := []testCase{
		{
			name: "success with nil used_at",
			input: &model.AdminPasswordReset{
				ID: fixedID, AdminUserID: adminUserID, OTPHash: resetOTPHash,
				Purpose: resetPurpose, AttemptCount: 0, ExpiresAt: expires, CreatedAt: created,
			},
			setup: func(mock sqlmock.Sqlmock, in *model.AdminPasswordReset) {
				mock.ExpectExec(wantInsertSQL).WithArgs(
					in.ID, in.AdminUserID, in.OTPHash, in.Purpose,
					in.AttemptCount, in.ExpiresAt, in.UsedAt, in.CreatedAt,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "success with used_at present",
			input: &model.AdminPasswordReset{
				ID: fixedID, AdminUserID: adminUserID, OTPHash: resetOTPHash,
				Purpose: resetPurpose, AttemptCount: 2, ExpiresAt: expires,
				UsedAt: &usedAt, CreatedAt: created,
			},
			setup: func(mock sqlmock.Sqlmock, in *model.AdminPasswordReset) {
				mock.ExpectExec(wantInsertSQL).WithArgs(
					in.ID, in.AdminUserID, in.OTPHash, in.Purpose,
					in.AttemptCount, in.ExpiresAt, in.UsedAt, in.CreatedAt,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: caseExecErr,
			input: &model.AdminPasswordReset{
				ID: fixedID, AdminUserID: adminUserID, OTPHash: resetOTPHash,
				Purpose: resetPurpose, ExpiresAt: expires, CreatedAt: created,
			},
			setup: func(mock sqlmock.Sqlmock, _ *model.AdminPasswordReset) {
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

			repo := adminPasswordReset.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
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
