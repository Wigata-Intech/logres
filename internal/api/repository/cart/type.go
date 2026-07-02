package cart

import (
	"database/sql"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
)

type cartRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func New(db *sql.DB, log *slog.Logger) repository.ICartRepository {
	return &cartRepository{
		db:  db,
		log: log,
	}
}
