package category

import (
	"database/sql"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
)

type categoryRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func New(db *sql.DB, log *slog.Logger) repository.ICategoryRepository {
	return &categoryRepository{
		db:  db,
		log: log,
	}
}
