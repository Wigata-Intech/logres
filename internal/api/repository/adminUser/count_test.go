package adminUser_test

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/api/repository/adminUser"
)

const wantCountSQL = `SELECT COUNT(*) FROM admin_users WHERE deleted_at IS NULL`

func TestCountAll(t *testing.T) {
	t.Parallel()

	errBoom := errors.New("query failed")

	type testCase struct {
		name      string
		setup     func(mock sqlmock.Sqlmock)
		wantErrIs error
		wantCount int
	}

	cases := []testCase{
		{
			name: "returns count",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"count"}).AddRow(int64(3))
				mock.ExpectQuery(wantCountSQL).WillReturnRows(rows)
			},
			wantCount: 3,
		},
		{
			name: "returns zero when empty",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"count"}).AddRow(int64(0))
				mock.ExpectQuery(wantCountSQL).WillReturnRows(rows)
			},
			wantCount: 0,
		},
		{
			name: "query error is wrapped",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(wantCountSQL).WillReturnError(errBoom)
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
			got, err := repo.CountAll(context.Background())

			if tc.wantErrIs != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.wantErrIs)
				assert.Zero(t, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.wantCount, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
