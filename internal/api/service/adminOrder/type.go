package adminOrder

import (
	"context"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
	"github.com/wigata-intech/logres/internal/api/service"
	"github.com/wigata-intech/logres/internal/shared/model"
)

type adminOrderService struct {
	logger          *slog.Logger
	cartRepository  repository.ICartRepository
	orderRepository repository.IOrderRepository
}

func New(
	logger *slog.Logger,
	cartRepository repository.ICartRepository,
	orderRepository repository.IOrderRepository,
) service.IAdminOrderService {
	return &adminOrderService{
		logger:          logger,
		cartRepository:  cartRepository,
		orderRepository: orderRepository,
	}
}

func (s *adminOrderService) List(ctx context.Context, req *model.AdminOrderListRequest) (*model.AdminOrderListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminOrderService) Detail(ctx context.Context, req *model.AdminOrderDetailRequest) (*model.AdminOrderResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminOrderService) Process(ctx context.Context, req *model.AdminOrderProcessRequest) (*model.AdminOrderResponse, error) {
	panic("not implemented") // TODO: Implement
}

// Cart Service
func (s *adminOrderService) CartList(ctx context.Context, req *model.AdminCartListRequest) (*model.AdminCartListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminOrderService) CartDetail(ctx context.Context, req *model.AdminCartDetailRequest) (*model.AdminCartResponse, error) {
	panic("not implemented") // TODO: Implement
}
