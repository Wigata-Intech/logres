package adminUser

import (
	"database/sql"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
)

type adminUserRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func New(db *sql.DB, log *slog.Logger) repository.IAdminUserRepository {
	return &adminUserRepository{
		db:  db,
		log: log,
	}
}
