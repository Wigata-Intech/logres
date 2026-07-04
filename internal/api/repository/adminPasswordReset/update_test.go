package adminPasswordReset_test

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

	"github.com/wigata-intech/logres/internal/api/repository/adminPasswordReset"
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

const wantIncrementAttemptSQL = `UPDATE admin_password_resets SET ` +
	`attempt_count = attempt_count + 1 ` +
	`WHERE id = ?`

const wantMarkUsedSQL = `UPDATE admin_password_resets SET ` +
	`used_at = ? ` +
	`WHERE id = ? AND used_at IS NULL`

const wantInvalidateAllSQL = `UPDATE admin_password_resets SET ` +
	`used_at = ? ` +
	`WHERE admin_user_id = ? AND purpose = ? AND used_at IS NULL`

func TestIncrementAttempt(t *testing.T) {
	t.Parallel()

	var (
		fixedID = uuid.MustParse("018f0000-0000-7000-8000-000000000001")
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
			name: "success increments one row",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantIncrementAttemptSQL).
					WithArgs(fixedID).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: caseExecErr,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantIncrementAttemptSQL).WillReturnError(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "rows-affected error is wrapped",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantIncrementAttemptSQL).
					WillReturnResult(sqlmock.NewErrorResult(errRows))
			},
			wantErr: errRows,
		},
		{
			name: "no rows affected maps to ErrNotFound",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantIncrementAttemptSQL).
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

			repo := adminPasswordReset.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			err = repo.IncrementAttempt(context.Background(), fixedID)

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

func TestMarkUsed(t *testing.T) {
	t.Parallel()

	var (
		fixedID = uuid.MustParse("018f0000-0000-7000-8000-000000000001")
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
			name: "success marks one row used",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantMarkUsedSQL).
					WithArgs(sqlmock.AnyArg(), fixedID).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: caseExecErr,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantMarkUsedSQL).WillReturnError(errBoom)
			},
			wantErr: errBoom,
		},
		{
			name: "rows-affected error is wrapped",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantMarkUsedSQL).
					WillReturnResult(sqlmock.NewErrorResult(errRows))
			},
			wantErr: errRows,
		},
		{
			name: "no rows affected maps to ErrNotFound",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantMarkUsedSQL).
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

			repo := adminPasswordReset.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			err = repo.MarkUsed(context.Background(), fixedID)

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

func TestInvalidateAllForUser(t *testing.T) {
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
			name: "success invalidates active resets",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantInvalidateAllSQL).
					WithArgs(sqlmock.AnyArg(), adminUserID, resetPurpose).
					WillReturnResult(sqlmock.NewResult(0, 2))
			},
		},
		{
			name: "success when nothing to invalidate",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantInvalidateAllSQL).
					WithArgs(sqlmock.AnyArg(), adminUserID, resetPurpose).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name: caseExecErr,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(wantInvalidateAllSQL).WillReturnError(errBoom)
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

			repo := adminPasswordReset.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			err = repo.InvalidateAllForUser(context.Background(), adminUserID, resetPurpose)

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
