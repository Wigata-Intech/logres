package api

import (
	"net/http"

	"github.com/wigata-intech/logres/internal/api/handler/admin"
	"github.com/wigata-intech/logres/internal/api/handler/cart"
	"github.com/wigata-intech/logres/internal/api/handler/order"
	"github.com/wigata-intech/logres/internal/api/handler/product"

	"github.com/wigata-intech/logres/config"
)

type Server struct {
	cfg            *config.APIServerConfig
	adminHandler   *admin.AdminHandler
	cartHandler    *cart.CartHandler
	orderHandler   *order.OrderHandler
	productHandler *product.ProductHandler
}

func New(opts ...OptionFunc) (http.Handler, error) {
	s := &Server{}

	for _, o := range opts {
		o(s)
	}

	mux := http.NewServeMux()

	// TODO: Need to group the routes

	// Admin API Group

	mux.HandleFunc("POST /api/v1/admin/user-setup", s.adminHandler.UserSetup)
	mux.HandleFunc("POST /api/v1/admin/auth/login", s.adminHandler.UserLogin)
	mux.HandleFunc("POST /api/v1/admin/auth/refresh-token", s.adminHandler.UserRefreshToken)
	mux.HandleFunc("POST /api/v1/admin/forgot-password", s.adminHandler.UserForgotPassword)
	mux.HandleFunc("POST /api/v1/admin/user/change-password", s.adminHandler.UserChangePassword)

	mux.HandleFunc("POST /api/v1/admin/file", s.adminHandler.FileUpload)
	mux.HandleFunc("GET /api/v1/admin/file", s.adminHandler.FileList)
	mux.HandleFunc("GET /api/v1/admin/file/{id}", s.adminHandler.FileDetail)
	mux.HandleFunc("PUT /api/v1/admin/file/{id}", s.adminHandler.FileUpdate)
	mux.HandleFunc("DELETE /api/v1/admin/file/{id}", s.adminHandler.FileDelete)

	mux.HandleFunc("POST /api/v1/admin/product", s.adminHandler.ProductCreate)
	mux.HandleFunc("GET /api/v1/admin/product", s.adminHandler.ProductList)
	mux.HandleFunc("GET /api/v1/admin/product/{id}", s.adminHandler.ProductDetail)
	mux.HandleFunc("PUT /api/v1/admin/product/{id}", s.adminHandler.ProductUpdate)
	mux.HandleFunc("DELETE /api/v1/admin/product/{id}", s.adminHandler.ProductDelete)

	mux.HandleFunc("POST /api/v1/admin/category", s.adminHandler.CategoryCreate)
	mux.HandleFunc("GET /api/v1/admin/category", s.adminHandler.CategoryList)
	mux.HandleFunc("PUT /api/v1/admin/category/{id}", s.adminHandler.CategoryUpdate)
	mux.HandleFunc("DELETE /api/v1/admin/category/{id}", s.adminHandler.CategoryDelete)

	mux.HandleFunc("POST /api/v1/admin/tag", s.adminHandler.TagCreate)
	mux.HandleFunc("GET /api/v1/admin/tag", s.adminHandler.TagList)
	mux.HandleFunc("PUT /api/v1/admin/tag/{id}", s.adminHandler.TagUpdate)
	mux.HandleFunc("DELETE /api/v1/admin/tag/{id}", s.adminHandler.TagDelete)

	mux.HandleFunc("GET /api/v1/admin/order", s.adminHandler.OrderList)
	mux.HandleFunc("GET /api/v1/admin/order/{id}", s.adminHandler.OrderDetail)
	mux.HandleFunc("PUT /api/v1/admin/order", s.adminHandler.OrderProcess)

	mux.HandleFunc("GET /api/v1/admin/cart", s.adminHandler.CartList)
	mux.HandleFunc("GET /api/v1/admin/cart/{id}", s.adminHandler.CartDetail)

	// Public API Group

	mux.HandleFunc("GET /api/v1/product", s.productHandler.List)
	mux.HandleFunc("GET /api/v1/product/{slug}", s.productHandler.Detail)

	// TODO: Need to think about the security
	mux.HandleFunc("POST /api/v1/cart", s.cartHandler.Insert)
	mux.HandleFunc("GET /api/v1/cart", s.cartHandler.List)
	mux.HandleFunc("PATCH /api/v1/cart/{id}", s.cartHandler.Patch)
	mux.HandleFunc("DELETE /api/v1/cart/{id}", s.cartHandler.Patch)

	// TODO: Add step for 2FA OTP
	mux.HandleFunc("POST /api/v1/checkout", s.orderHandler.Checkout)
	mux.HandleFunc("POST /api/v1/check-order", s.orderHandler.Check)

	return mux, nil
}
