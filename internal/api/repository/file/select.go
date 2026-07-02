package file

import (
	"context"

	"github.com/wigata-intech/logres/internal/api/model"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"

	"github.com/google/uuid"
)

func (r *fileRepository) List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.File, *sharedModel.Pagination, error) {
	panic("not implemented") // TODO: Implement
}

func (r *fileRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.File, error) {
	panic("not implemented") // TODO: Implement
}
