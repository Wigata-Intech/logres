package product

import (
	"context"

	"github.com/wigata-intech/logres/internal/api/model"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"

	"github.com/google/uuid"
)

func (r *productRepository) List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.Product, *sharedModel.Pagination, error) {
	panic("not implemented") // TODO: Implement
}

func (r *productRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (r *productRepository) GetBySlug(ctx context.Context, slug string) (*model.Product, error) {
	panic("not implemented") // TODO: Implement
}
