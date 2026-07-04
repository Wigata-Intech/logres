package adminPasswordReset

import (
	"database/sql"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
)

const (
	insertAdminPasswordReset = `INSERT INTO admin_password_resets ` +
		`(id, admin_user_id, otp_hash, purpose, attempt_count, expires_at, used_at, created_at) ` +
		`VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	selectActiveAdminPasswordResetByUser = `SELECT ` +
		`id, admin_user_id, otp_hash, purpose, attempt_count, expires_at, used_at, created_at ` +
		`FROM admin_password_resets ` +
		`WHERE admin_user_id = ? AND purpose = ? AND used_at IS NULL ` +
		`ORDER BY created_at DESC LIMIT 1`

	incrementAdminPasswordResetAttempt = `UPDATE admin_password_resets SET ` +
		`attempt_count = attempt_count + 1 ` +
		`WHERE id = ?`

	markAdminPasswordResetUsed = `UPDATE admin_password_resets SET ` +
		`used_at = ? ` +
		`WHERE id = ? AND used_at IS NULL`

	invalidateAdminPasswordResetsForUser = `UPDATE admin_password_resets SET ` +
		`used_at = ? ` +
		`WHERE admin_user_id = ? AND purpose = ? AND used_at IS NULL`
)

type adminPasswordResetRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func New(db *sql.DB, log *slog.Logger) repository.IAdminPasswordResetRepository {
	return &adminPasswordResetRepository{db: db, log: log}
}
