package cart

import (
	"context"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
	"github.com/wigata-intech/logres/internal/api/service"
	"github.com/wigata-intech/logres/internal/shared/model"
)

type cartService struct {
	logger            *slog.Logger
	cartRepository    repository.ICartRepository
	productRepository repository.IProductRepository
}

func New(
	logger *slog.Logger,
	cartRepository repository.ICartRepository,
	productRepository repository.IProductRepository,
) service.ICartService {
	return &cartService{
		logger:            logger,
		cartRepository:    cartRepository,
		productRepository: productRepository,
	}
}

func (s *cartService) Insert(ctx context.Context, req *model.CartInsertRequest) (*model.CartResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *cartService) List(ctx context.Context, req *model.CartListRequest) (*model.CartListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *cartService) Patch(ctx context.Context, req *model.CartPatchRequest) (*model.CartResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *cartService) Delete(ctx context.Context, req *model.CartDeleteRequest) error {
	panic("not implemented") // TODO: Implement
}
