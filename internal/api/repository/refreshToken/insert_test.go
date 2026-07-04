package refreshToken_test

import (
	"context"
	"encoding/json"
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
	"github.com/wigata-intech/logres/internal/api/repository/refreshToken"
)

const caseExecErr = "exec error is wrapped"

// wantInsertSQL is matched exactly (QueryMatcherEqual) so the test guards the
// statement's column list and order, not just the bound args.
const wantInsertSQL = `INSERT INTO admin_refresh_tokens ` +
	`(id, admin_user_id, token_hash, family_id, replaced_by, expires_at, revoked_at, created_at, metadata, last_used_at) ` +
	`VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

const tokenHash = "deadbeefcafe"

const (
	tokenUserAgent = "Mozilla/5.0"
	tokenIPAddress = "203.0.113.10"
	tokenDeviceLbl = "Chrome on macOS"
)

func marshalMetadata(t *testing.T, m model.RefreshTokenMetadata) string {
	t.Helper()
	b, err := json.Marshal(m)
	require.NoError(t, err)
	return string(b)
}

func TestCreate(t *testing.T) {
	t.Parallel()

	var (
		fixedID     = uuid.MustParse("018f0000-0000-7000-8000-000000000001")
		adminUserID = uuid.MustParse("018f0000-0000-7000-8000-000000000002")
		familyID    = uuid.MustParse("018f0000-0000-7000-8000-000000000003")
		replacedBy  = uuid.MustParse("018f0000-0000-7000-8000-000000000004")
		expires     = time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
		created     = time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
		revoked     = time.Date(2026, 1, 3, 0, 0, 0, 0, time.UTC)
		deviceLabel = tokenDeviceLbl
		errBoom     = errors.New("exec failed")
	)

	deviceNoLabel := model.RefreshTokenDevice{UserAgent: tokenUserAgent, IPAddress: tokenIPAddress}
	deviceWithLabel := model.RefreshTokenDevice{UserAgent: tokenUserAgent, IPAddress: tokenIPAddress, DeviceLabel: &deviceLabel}
	metadataNoLabel := model.RefreshTokenMetadata{Device: deviceNoLabel}
	metadataWithLabel := model.RefreshTokenMetadata{Device: deviceWithLabel}

	type testCase struct {
		name    string
		input   *model.RefreshToken
		setup   func(mock sqlmock.Sqlmock, in *model.RefreshToken)
		wantErr error
	}

	cases := []testCase{
		{
			name: "success with nullable fields absent",
			input: &model.RefreshToken{
				ID: fixedID, AdminUserID: adminUserID, TokenHash: tokenHash,
				FamilyID: familyID, ExpiresAt: expires, CreatedAt: created,
				Metadata: metadataNoLabel, LastUsedAt: created,
			},
			setup: func(mock sqlmock.Sqlmock, in *model.RefreshToken) {
				mock.ExpectExec(wantInsertSQL).WithArgs(
					in.ID, in.AdminUserID, in.TokenHash, in.FamilyID,
					in.ReplacedBy, in.ExpiresAt, in.RevokedAt, in.CreatedAt,
					marshalMetadata(t, in.Metadata), in.LastUsedAt,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "success with nullable fields present",
			input: &model.RefreshToken{
				ID: fixedID, AdminUserID: adminUserID, TokenHash: tokenHash,
				FamilyID: familyID, ReplacedBy: &replacedBy, ExpiresAt: expires,
				RevokedAt: &revoked, CreatedAt: created,
				Metadata: metadataWithLabel, LastUsedAt: created,
			},
			setup: func(mock sqlmock.Sqlmock, in *model.RefreshToken) {
				mock.ExpectExec(wantInsertSQL).WithArgs(
					in.ID, in.AdminUserID, in.TokenHash, in.FamilyID,
					in.ReplacedBy, in.ExpiresAt, in.RevokedAt, in.CreatedAt,
					marshalMetadata(t, in.Metadata), in.LastUsedAt,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: caseExecErr,
			input: &model.RefreshToken{
				ID: fixedID, AdminUserID: adminUserID, TokenHash: tokenHash,
				FamilyID: familyID, ExpiresAt: expires, CreatedAt: created,
				Metadata: metadataNoLabel, LastUsedAt: created,
			},
			setup: func(mock sqlmock.Sqlmock, _ *model.RefreshToken) {
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

			repo := refreshToken.New(db, slog.New(slog.NewTextHandler(io.Discard, nil)))
			err = repo.Insert(context.Background(), tc.input)

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
