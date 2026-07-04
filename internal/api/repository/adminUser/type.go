package adminUser

import (
	"database/sql"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
)

const (
	insertAdminUser = `INSERT INTO admin_users ` +
		`(id, full_name, email, password, status, last_login_at, created_at, updated_at, deleted_at) ` +
		`VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	selectAdminUserByEmail = `SELECT ` +
		`id, full_name, email, password, status, last_login_at, created_at, updated_at, deleted_at ` +
		`FROM admin_users WHERE email = ? AND deleted_at IS NULL`

	selectAdminUserByID = `SELECT ` +
		`id, full_name, email, password, status, last_login_at, created_at, updated_at, deleted_at ` +
		`FROM admin_users WHERE id = ? AND deleted_at IS NULL`

	updateAdminUser = `UPDATE admin_users SET ` +
		`full_name = ?, email = ?, password = ?, status = ?, ` +
		`last_login_at = ?, updated_at = ?, deleted_at = ? ` +
		`WHERE id = ?`

	countAdminUsers = `SELECT COUNT(*) FROM admin_users WHERE deleted_at IS NULL`
)

type adminUserRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func New(db *sql.DB, log *slog.Logger) repository.IAdminUserRepository {
	return &adminUserRepository{db: db, log: log}
}
