package adminRefreshToken

import (
	"database/sql"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
)

const (
	insertAdminRefreshToken = `INSERT INTO admin_refresh_tokens ` +
		`(id, admin_user_id, token_hash, family_id, replaced_by, expires_at, revoked_at, created_at, metadata, last_used_at) ` +
		`VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	selectAdminRefreshTokenByHash = `SELECT ` +
		`id, admin_user_id, token_hash, family_id, replaced_by, expires_at, revoked_at, created_at, metadata, last_used_at ` +
		`FROM admin_refresh_tokens WHERE token_hash = ?`

	rotateAdminRefreshToken = `UPDATE admin_refresh_tokens SET ` +
		`replaced_by = ?, revoked_at = ? ` +
		`WHERE id = ? AND revoked_at IS NULL`

	revokeAdminRefreshTokenFamily = `UPDATE admin_refresh_tokens SET ` +
		`revoked_at = ? ` +
		`WHERE family_id = ? AND revoked_at IS NULL`

	revokeAllAdminRefreshTokensForUser = `UPDATE admin_refresh_tokens SET ` +
		`revoked_at = ? ` +
		`WHERE admin_user_id = ? AND revoked_at IS NULL`
)

type adminRefreshTokenRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func New(db *sql.DB, log *slog.Logger) repository.IAdminRefreshTokenRepository {
	return &adminRefreshTokenRepository{db: db, log: log}
}
