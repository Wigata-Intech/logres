package product

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/api/repository"

	"github.com/google/uuid"
)

type productRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func New(db *sql.DB, log *slog.Logger) repository.IProductRepository {
	return &productRepository{
		db:  db,
		log: log,
	}
}

// Files
func (r *productRepository) FileAssign(ctx context.Context, productFile *model.ProductFile) error {
	panic("not implemented") // TODO: Implement
}

func (r *productRepository) FileUnassign(ctx context.Context, productFile *model.ProductFile) error {
	panic("not implemented") // TODO: Implement
}

func (r *productRepository) FileListByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductFile, error) {
	panic("not implemented") // TODO: Implement
}

// Category and Tags
func (r *productRepository) CategoryAssign(ctx context.Context, productCategory *model.ProductCategory) error {
	panic("not implemented") // TODO: Implement
}

func (r *productRepository) CategoryUnassign(ctx context.Context, productCategory *model.ProductCategory) error {
	panic("not implemented") // TODO: Implement
}

func (r *productRepository) CategoryListByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductCategory, error) {
	panic("not implemented") // TODO: Implement
}

func (r *productRepository) TagAssign(ctx context.Context, productTag *model.ProductTag) error {
	panic("not implemented") // TODO: Implement
}

func (r *productRepository) TagUnassign(ctx context.Context, productTag *model.ProductTag) error {
	panic("not implemented") // TODO: Implement
}

func (r *productRepository) TagListByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductTag, error) {
	panic("not implemented") // TODO: Implement
}
