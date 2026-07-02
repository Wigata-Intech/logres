package order

import (
	"context"

	"github.com/wigata-intech/logres/internal/api/model"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"

	"github.com/google/uuid"
)

func (r *orderRepository) List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.Order, *sharedModel.Pagination, error) {
	panic("not implemented") // TODO: Implement
}

func (r *orderRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	panic("not implemented") // TODO: Implement
}
