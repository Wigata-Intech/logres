package service

import (
	"context"

	"github.com/wigata-intech/logres/internal/shared/model"
)

type IAdminUserService interface {
	Setup(ctx context.Context, req *model.AdminSetupRequest) (*model.AdminSetupResponse, error)
	Login(ctx context.Context, req *model.AdminLoginRequest) (*model.AdminTokenResponse, error)
	RefreshToken(ctx context.Context, req *model.AdminRefreshTokenRequest) (*model.AdminTokenResponse, error)
	ForgotPassword(ctx context.Context, req *model.AdminForgotPasswordRequest) (*model.AdminForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, req *model.AdminResetPasswordRequest) (*model.AdminResetPasswordResponse, error)
	ChangePassword(ctx context.Context, req *model.AdminChangePasswordRequest) (*model.AdminChangePasswordResponse, error)
}

type IAdminFileService interface {
	Upload(ctx context.Context, req *model.AdminFileUploadRequest) (*model.AdminFileResponse, error)
	List(ctx context.Context, req *model.AdminFileListRequest) (*model.AdminFileListResponse, error)
	Detail(ctx context.Context, req *model.AdminFileDetailRequest) (*model.AdminFileResponse, error)
	Update(ctx context.Context, req *model.AdminFileUpdateRequest) (*model.AdminFileResponse, error)
	Delete(ctx context.Context, req *model.AdminFileDeleteRequest) error
}

type IAdminProductService interface {
	Create(ctx context.Context, req *model.AdminProductCreateRequest) (*model.AdminProductResponse, error)
	List(ctx context.Context, req *model.AdminProductListRequest) (*model.AdminProductListResponse, error)
	Detail(ctx context.Context, req *model.AdminProductDetailRequest) (*model.AdminProductResponse, error)
	Update(ctx context.Context, req *model.AdminProductUpdateRequest) (*model.AdminProductResponse, error)
	Delete(ctx context.Context, req *model.AdminProductDeleteRequest) error

	// Category and Tags
	CategoryCreate(ctx context.Context, req *model.AdminCategoryCreateRequest) (*model.AdminCategoryResponse, error)
	CategoryList(ctx context.Context, req *model.AdminCategoryListRequest) (*model.AdminCategoryListResponse, error)
	CategoryUpdate(ctx context.Context, req *model.AdminCategoryUpdateRequest) (*model.AdminCategoryResponse, error)
	CategoryDelete(ctx context.Context, req *model.AdminCategoryDeleteRequest) error
	TagCreate(ctx context.Context, req *model.AdminTagCreateRequest) (*model.AdminTagResponse, error)
	TagList(ctx context.Context, req *model.AdminTagListRequest) (*model.AdminTagListResponse, error)
	TagUpdate(ctx context.Context, req *model.AdminTagUpdateRequest) (*model.AdminTagResponse, error)
	TagDelete(ctx context.Context, req *model.AdminTagDeleteRequest) error
}

type IAdminOrderService interface {
	List(ctx context.Context, req *model.AdminOrderListRequest) (*model.AdminOrderListResponse, error)
	Detail(ctx context.Context, req *model.AdminOrderDetailRequest) (*model.AdminOrderResponse, error)
	Process(ctx context.Context, req *model.AdminOrderProcessRequest) (*model.AdminOrderResponse, error)

	// Cart Service
	CartList(ctx context.Context, req *model.AdminCartListRequest) (*model.AdminCartListResponse, error)
	CartDetail(ctx context.Context, req *model.AdminCartDetailRequest) (*model.AdminCartResponse, error)
}

type IProductService interface {
	List(ctx context.Context, req *model.ProductListRequest) (*model.ProductListResponse, error)
	Detail(Ctx context.Context, req *model.ProductDetailRequest) (*model.ProductResponse, error)
}

type ICartService interface {
	Insert(ctx context.Context, req *model.CartInsertRequest) (*model.CartResponse, error)
	List(ctx context.Context, req *model.CartListRequest) (*model.CartListResponse, error)
	Patch(ctx context.Context, req *model.CartPatchRequest) (*model.CartResponse, error)
	Delete(ctx context.Context, req *model.CartDeleteRequest) error
}

type IOrderService interface {
	Checkout(ctx context.Context, req *model.OrderCheckoutRequest) (*model.OrderCheckoutResponse, error)
	Check(ctx context.Context, req *model.OrderCheckRequest) (*model.OrderResponse, error)
}
