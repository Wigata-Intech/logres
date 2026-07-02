package order

import (
	"context"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
	"github.com/wigata-intech/logres/internal/api/service"
	"github.com/wigata-intech/logres/internal/shared/model"
)

type orderService struct {
	logger            *slog.Logger
	cartRepository    repository.ICartRepository
	orderRepository   repository.IOrderRepository
	productRepository repository.IProductRepository
}

func New(
	logger *slog.Logger,
	cartRepository repository.ICartRepository,
	orderRepository repository.IOrderRepository,
	productRepository repository.IProductRepository,
) service.IOrderService {
	return &orderService{
		logger:            logger,
		cartRepository:    cartRepository,
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (s *orderService) Checkout(ctx context.Context, req *model.OrderCheckoutRequest) (*model.OrderCheckoutResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *orderService) Check(ctx context.Context, req *model.OrderCheckRequest) (*model.OrderResponse, error) {
	panic("not implemented") // TODO: Implement
}
