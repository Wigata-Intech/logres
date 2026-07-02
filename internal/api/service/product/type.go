package product

import (
	"context"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
	"github.com/wigata-intech/logres/internal/api/service"
	"github.com/wigata-intech/logres/internal/shared/model"
)

type productService struct {
	logger            *slog.Logger
	productRepository repository.IProductRepository
}

func New(
	logger *slog.Logger,
	productRepository repository.IProductRepository,
) service.IProductService {
	return &productService{
		logger:            logger,
		productRepository: productRepository,
	}
}

func (s *productService) List(ctx context.Context, req *model.ProductListRequest) (*model.ProductListResponse, error) {
	panic("not implement") // TODO: Implement
}

func (s *productService) Detail(Ctx context.Context, req *model.ProductDetailRequest) (*model.ProductResponse, error) {
	panic("not implement") // TODO: Implement
}
