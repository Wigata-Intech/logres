package category

import (
	"context"

	"github.com/wigata-intech/logres/internal/api/model"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"
)

func (r *categoryRepository) List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.Category, *sharedModel.Pagination, error) {
	panic("not implemented") // TODO: Implement
}

func (r *categoryRepository) GetBySlug(ctx context.Context, slug string) (*model.Category, error) {
	panic("not implemented") // TODO: Implement
}
