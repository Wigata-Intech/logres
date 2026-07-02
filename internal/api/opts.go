package api

import (
	"github.com/wigata-intech/logres/internal/api/handler/admin"
	"github.com/wigata-intech/logres/internal/api/handler/cart"
	"github.com/wigata-intech/logres/internal/api/handler/order"
	"github.com/wigata-intech/logres/internal/api/handler/product"
)

type OptionFunc func(*Server)

func WithAdminHandler(adminHandler *admin.AdminHandler) OptionFunc {
	return func(s *Server) { s.adminHandler = adminHandler }
}

func WithCartHandler(cartHandler *cart.CartHandler) OptionFunc {
	return func(s *Server) { s.cartHandler = cartHandler }
}

func WithOrderHandler(orderHandler *order.OrderHandler) OptionFunc {
	return func(s *Server) { s.orderHandler = orderHandler }
}

func WithProductHandler(productHandler *product.ProductHandler) OptionFunc {
	return func(s *Server) { s.productHandler = productHandler }
}
