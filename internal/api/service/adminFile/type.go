package adminFile

import (
	"context"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
	"github.com/wigata-intech/logres/internal/api/service"
	"github.com/wigata-intech/logres/internal/shared/model"
)

type adminFileService struct {
	logger         *slog.Logger
	fileRepository repository.IFileRepository
}

func New(
	logger *slog.Logger,
	fileRepository repository.IFileRepository,
) service.IAdminFileService {
	return &adminFileService{
		logger:         logger,
		fileRepository: fileRepository,
	}
}

func (s *adminFileService) Upload(ctx context.Context, req *model.AdminFileUploadRequest) (*model.AdminFileResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminFileService) List(ctx context.Context, req *model.AdminFileListRequest) (*model.AdminFileListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminFileService) Detail(ctx context.Context, req *model.AdminFileDetailRequest) (*model.AdminFileResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminFileService) Update(ctx context.Context, req *model.AdminFileUpdateRequest) (*model.AdminFileResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminFileService) Delete(ctx context.Context, req *model.AdminFileDeleteRequest) error {
	panic("not implemented") // TODO: Implement
}
