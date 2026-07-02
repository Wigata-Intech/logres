package cart

import (
	"context"

	"github.com/wigata-intech/logres/internal/api/model"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"

	"github.com/google/uuid"
)

func (r *cartRepository) List(ctx context.Context, pagination *sharedModel.Pagination, identifierID uuid.UUID) ([]model.Cart, *sharedModel.Pagination, error) {
	panic("not implemented") // TODO: Implement
}

func (r *cartRepository) GetByIdentifierID(ctx context.Context, identifierID uuid.UUID) (*model.Cart, error) {
	panic("not implemented") // TODO: Implement
}
