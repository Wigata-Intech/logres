package adminRefreshToken_test

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/api/repository/adminRefreshToken"
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

const wantRotateSQL = `UPDATE admin_refresh_tokens SET ` +
	`replaced_by = ?, revoked_at = ? ` +
	`WHERE id = ? AND revoked_at IS NULL`

const wantRevokeFamilySQL = `UPDATE admin_refresh_tokens SET ` +
	`revoked_at = ? ` +
	`WHERE family_id = ? AND revoked_at IS NULL`

const wantRevokeAllForUserSQL = `UPDATE admin_refresh_tokens SET ` +
	`revoked_at = ? ` +
	`WHERE admin_user_id = ? AND revoked_at IS NULL`

func TestRotate(t *testing.T) {
	t.Parallel()

	var (
		oldID   = uuid.MustParse("018f0000-0000-7000-8000-000000000001")
		newID   = uuid.MustParse("018f0000-0000-7000-8000-000000000002")
		errBoom = errors.New("exec failed")
		errRows = errors.New("rows affected failed")
	)

	type testCase struct {
		name    string
		setup   func(mock sqlmock.Sqlmock)
		wantErr error
	}

	cases := []testCase{
		{
			name: "success rotates one row",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRotateSQL).
					WithArgs(newID, sqlmock.AnyArg(), oldID).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: caseExecErr,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRotateSQL).WillReturnError(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "rows-affected error is wrapped",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRotateSQL).
					WillReturnResult(sqlmock.NewErrorResult(errRows))
			},
			wantErr: errRows,
		},
		{
			name: "no rows affected maps to ErrNotFound",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRotateSQL).
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

			tc.setup(mock)

			repo := adminRefreshToken.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			err = repo.Rotate(context.Background(), oldID, newID)

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

func TestRevokeFamily(t *testing.T) {
	t.Parallel()

	var (
		familyID = uuid.MustParse("018f0000-0000-7000-8000-000000000003")
		errBoom  = errors.New("exec failed")
	)

	type testCase struct {
		name    string
		setup   func(mock sqlmock.Sqlmock)
		wantErr error
	}

	cases := []testCase{
		{
			name: "success revokes family",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRevokeFamilySQL).
					WithArgs(sqlmock.AnyArg(), familyID).
					WillReturnResult(sqlmock.NewResult(0, 2))
			},
		},
		{
			name: "success when family already revoked",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRevokeFamilySQL).
					WithArgs(sqlmock.AnyArg(), familyID).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name: caseExecErr,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRevokeFamilySQL).WillReturnError(errBoom)
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

			tc.setup(mock)

			repo := adminRefreshToken.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			err = repo.RevokeFamily(context.Background(), familyID)

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

func TestRevokeAllForUser(t *testing.T) {
	t.Parallel()

	var (
		adminUserID = uuid.MustParse("018f0000-0000-7000-8000-000000000002")
		errBoom     = errors.New("exec failed")
	)

	type testCase struct {
		name    string
		setup   func(mock sqlmock.Sqlmock)
		wantErr error
	}

	cases := []testCase{
		{
			name: "success revokes all sessions",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRevokeAllForUserSQL).
					WithArgs(sqlmock.AnyArg(), adminUserID).
					WillReturnResult(sqlmock.NewResult(0, 3))
			},
		},
		{
			name: "success when nothing to revoke",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRevokeAllForUserSQL).
					WithArgs(sqlmock.AnyArg(), adminUserID).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name: caseExecErr,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantRevokeAllForUserSQL).WillReturnError(errBoom)
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

			tc.setup(mock)

			repo := adminRefreshToken.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			err = repo.RevokeAllForUser(context.Background(), adminUserID)

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
