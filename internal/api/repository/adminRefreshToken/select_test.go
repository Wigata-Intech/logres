package adminRefreshToken_test

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
	"github.com/wigata-intech/logres/internal/api/repository/adminRefreshToken"
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

const wantSelectByHashSQL = `SELECT ` +
	`id, admin_user_id, token_hash, family_id, replaced_by, expires_at, revoked_at, created_at, metadata, last_used_at ` +
	`FROM admin_refresh_tokens WHERE token_hash = ?`

var adminRefreshTokenColumns = []string{
	"id", "admin_user_id", "token_hash", "family_id",
	"replaced_by", "expires_at", "revoked_at", "created_at",
	"metadata", "last_used_at",
}

func TestGetByHash(t *testing.T) {
	t.Parallel()

	const (
		idText          = "018f0000-0000-7000-8000-000000000001"
		adminUserIDText = "018f0000-0000-7000-8000-000000000002"
		familyIDText    = "018f0000-0000-7000-8000-000000000003"
		replacedByText  = "018f0000-0000-7000-8000-000000000004"
	)
	var (
		fixedID     = uuid.MustParse(idText)
		adminUserID = uuid.MustParse(adminUserIDText)
		familyID    = uuid.MustParse(familyIDText)
		replacedBy  = uuid.MustParse(replacedByText)
		expires     = time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
		created     = time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
		revoked     = time.Date(2026, 1, 3, 0, 0, 0, 0, time.UTC)
		lastUsed    = time.Date(2026, 1, 2, 4, 0, 0, 0, time.UTC)
		deviceLabel = tokenDeviceLbl
		errBoom     = errors.New("query failed")
	)

	deviceNoLabel := model.AdminRefreshTokenDevice{UserAgent: tokenUserAgent, IPAddress: tokenIPAddress}
	deviceWithLabel := model.AdminRefreshTokenDevice{UserAgent: tokenUserAgent, IPAddress: tokenIPAddress, DeviceLabel: &deviceLabel}
	metadataNoLabel := model.AdminRefreshTokenMetadata{Device: deviceNoLabel}
	metadataWithLabel := model.AdminRefreshTokenMetadata{Device: deviceWithLabel}

	type testCase struct {
		name       string
		setup      func(mock sqlmock.Sqlmock)
		wantErrIs  error
		wantErr    bool
		wantResult *model.AdminRefreshToken
	}

	cases := []testCase{
		{
			name: "found active token with no device label",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows(adminRefreshTokenColumns).AddRow(
					idText, adminUserIDText, tokenHash, familyIDText,
					nil, expires, nil, created,
					marshalMetadata(t, metadataNoLabel), lastUsed,
				)
				mock.ExpectQuery(wantSelectByHashSQL).WithArgs(tokenHash).WillReturnRows(rows)
			},
			wantResult: &model.AdminRefreshToken{
				ID: fixedID, AdminUserID: adminUserID, TokenHash: tokenHash,
				FamilyID: familyID, ExpiresAt: expires, CreatedAt: created,
				Metadata: metadataNoLabel, LastUsedAt: lastUsed,
			},
		},
		{
			name: "found rotated and revoked token with device label",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows(adminRefreshTokenColumns).AddRow(
					idText, adminUserIDText, tokenHash, familyIDText,
					replacedByText, expires, revoked, created,
					marshalMetadata(t, metadataWithLabel), lastUsed,
				)
				mock.ExpectQuery(wantSelectByHashSQL).WithArgs(tokenHash).WillReturnRows(rows)
			},
			wantResult: &model.AdminRefreshToken{
				ID: fixedID, AdminUserID: adminUserID, TokenHash: tokenHash,
				FamilyID: familyID, ReplacedBy: &replacedBy, ExpiresAt: expires,
				RevokedAt: &revoked, CreatedAt: created,
				Metadata: metadataWithLabel, LastUsedAt: lastUsed,
			},
		},
		{
			name: "not found maps to ErrNotFound",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(wantSelectByHashSQL).WithArgs(tokenHash).WillReturnError(sql.ErrNoRows)
			},
			wantErrIs: dbx.ErrNotFound,
		},
		{
			name: "scan error is wrapped",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(wantSelectByHashSQL).WithArgs(tokenHash).WillReturnError(errBoom)
			},
			wantErrIs: errBoom,
		},
		{
			name: "malformed metadata JSON is wrapped",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows(adminRefreshTokenColumns).AddRow(
					idText, adminUserIDText, tokenHash, familyIDText,
					nil, expires, nil, created,
					`{not-json`, lastUsed,
				)
				mock.ExpectQuery(wantSelectByHashSQL).WithArgs(tokenHash).WillReturnRows(rows)
			},
			wantErr: true,
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
			got, err := repo.GetByHash(context.Background(), tokenHash)

			switch {
			case tc.wantErrIs != nil:
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.wantErrIs)
				assert.Nil(t, got)
			case tc.wantErr:
				require.Error(t, err)
				assert.Nil(t, got)
			default:
				require.NoError(t, err)
				require.NotNil(t, got)
				assertAdminRefreshTokenEqual(t, tc.wantResult, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func assertAdminRefreshTokenEqual(t *testing.T, want, got *model.AdminRefreshToken) {
	t.Helper()
	assert.Equal(t, want.ID, got.ID)
	assert.Equal(t, want.AdminUserID, got.AdminUserID)
	assert.Equal(t, want.TokenHash, got.TokenHash)
	assert.Equal(t, want.FamilyID, got.FamilyID)
	assert.Equal(t, want.ReplacedBy, got.ReplacedBy)
	assert.True(t, want.ExpiresAt.Equal(got.ExpiresAt))
	assert.True(t, want.CreatedAt.Equal(got.CreatedAt))
	assert.Equal(t, want.Metadata, got.Metadata)
	assert.True(t, want.LastUsedAt.Equal(got.LastUsedAt))
	if want.RevokedAt == nil {
		assert.Nil(t, got.RevokedAt)
	} else {
		require.NotNil(t, got.RevokedAt)
		assert.True(t, want.RevokedAt.Equal(*got.RevokedAt))
	}
}
