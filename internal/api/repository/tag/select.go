package tag

import (
	"context"

	"github.com/wigata-intech/logres/internal/api/model"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"

	"github.com/google/uuid"
)

func (r *tagRepository) List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.Tag, *sharedModel.Pagination, error) {
	panic("not implemented") // TODO: Implement
}

func (r *tagRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Tag, error) {
	panic("not implemented") // TODO: Implement
}
