package tag

import (
	"database/sql"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
)

type tagRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func New(db *sql.DB, log *slog.Logger) repository.ITagRepository {
	return &tagRepository{
		db:  db,
		log: log,
	}
}
