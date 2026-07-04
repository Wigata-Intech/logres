package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Admin User

type (
	AdminSetupRequest struct {
		FullName string `json:"full_name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	AdminLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	AdminRefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	AdminForgotPasswordRequest struct {
		Email string `json:"email" validate:"required,email"`
	}

	AdminChangePasswordRequest struct {
		UserID      uuid.UUID `json:"user_id" validate:"required"`
		OldPassword string    `json:"old_password" validate:"required"`
		NewPassword string    `json:"new_password" validate:"required"`
	}

	AdminResetPasswordRequest struct {
		Email       string `json:"email" validate:"required,email"`
		OTP         string `json:"otp" validate:"required,len=8,numeric"`
		NewPassword string `json:"new_password" validate:"required"`
	}
)

// Admin File Management Service

type (
	AdminFileUploadRequest struct {
		FileName     string `json:"file_name" validate:"required"`
		FilePath     string `json:"file_path" validate:"required"`
		FileMimeType string `json:"file_mime_type" validate:"required,valMimeType"`
		FileBlob     string `json:"file_blob" validate:"required"` // TODO: For now, it's required. In the future depends
		IsPublic     bool   `json:"is_public" validate:"-"`
	}

	AdminFileListRequest struct {
		Token  string `json:"token" validate:"-"`
		Limit  uint16 `json:"limit" validate:"required"`
		Search string `json:"search" validate:"-"`
	}

	AdminFileDetailRequest struct {
		FileID uuid.UUID `json:"file_id" validate:"required"`
	}

	AdminFileUpdateRequest struct {
		FileID   uuid.UUID `json:"file_id" validate:"required"`
		FileName string    `json:"file_name" validate:"-"`
		FilePath string    `json:"file_path" validate:"-"`
		IsPublic bool      `json:"is_public" validate:"-"`
	}

	AdminFileDeleteRequest struct {
		FileID uuid.UUID `json:"file_id" validate:"required"`
	}
)

// Admin Product Service

type (
	AdminProductCreateRequest struct {
		ProductName        string          `json:"product_name" validate:"required"`
		ProductDescription string          `json:"product_description" validate:"required"`
		ProductPrice       decimal.Decimal `json:"product_price" validate:"required"`
		// TODO: Add for Files, Category and Tags. Files and Tags can have more than one
	}

	AdminProductListRequest struct {
		Token         string `json:"token" validate:"-"`
		Limit         uint16 `json:"limit" validate:"required"`
		Search        string `json:"search" validate:"-"`
		ProductStatus uint8  `json:"product_status" validate:"valProductStatus"`
	}

	AdminProductDetailRequest struct {
		ProductID uuid.UUID `json:"product_id" validate:"required"`
	}

	AdminProductUpdateRequest struct {
		ProductName        string          `json:"product_name" validate:"required"`
		ProductDescription string          `json:"product_description" validate:"required"`
		ProductPrice       decimal.Decimal `json:"product_price" validate:"required"`
		// TODO: Add for Files, Category and Tags. Files and Tags can have more than one
	}

	AdminProductDeleteRequest struct {
		ProductID uuid.UUID `json:"product_id" validate:"required"`
	}

	// Category and Tags
	AdminCategoryCreateRequest struct {
		CategoryName string `json:"category_name" validate:"required"`
	}

	AdminCategoryListRequest struct {
		Token  string `json:"token" validate:"-"`
		Limit  uint16 `json:"limit" validate:"required"`
		Search string `json:"search" validate:"-"`
	}

	AdminCategoryUpdateRequest struct {
		CategoryName string `json:"category_name" validate:"required"`
	}

	AdminCategoryDeleteRequest struct {
		CategoryID uuid.UUID `json:"category_id" validate:"required"`
	}

	AdminTagCreateRequest struct {
		TagName string `json:"tag_name" validate:"required"`
	}

	AdminTagListRequest struct {
		Token  string `json:"token" validate:"-"`
		Limit  uint16 `json:"limit" validate:"required"`
		Search string `json:"search" validate:"-"`
	}

	AdminTagUpdateRequest struct {
		TagName string `json:"tag_name" validate:"required"`
	}

	AdminTagDeleteRequest struct {
		TagID uuid.UUID `json:"tag_id" validate:"required"`
	}
)

// Admin Order Service

type (
	AdminOrderListRequest struct {
		Token       string `json:"token" validate:"-"`
		Limit       uint16 `json:"limit" validate:"required"`
		Search      string `json:"search" validate:"-"`
		OrderStatus string `json:"order_status" validate:"valOrderStatus"`
	}

	AdminOrderDetailRequest struct {
		OrderID uuid.UUID `json:"order_id" validate:"required"`
	}

	// Handling for order processing
	// Such as payment confirmation, shipping, and delivery etc.
	//
	// TODO: Every status have different payload?
	// Need to make the FSM first, and we know what is the needs
	AdminOrderProcessRequest struct {
	}

	// Cart Service
	AdminCartListRequest struct {
		Token  string `json:"token" validate:"-"`
		Limit  uint16 `json:"limit" validate:"required"`
		Search string `json:"search" validate:"-"`
	}

	AdminCartDetailRequest struct {
		CartIdentifierID uuid.UUID `json:"cart_identifier_id" validate:"required"`
	}
)

// Product Service

type (
	ProductListRequest struct {
		Token  string `json:"token" validate:"-"`
		Limit  uint16 `json:"limit" validate:"required"`
		Search string `json:"search" validate:"-"`
	}

	ProductDetailRequest struct {
		ProductSlug string `json:"product_slug" valdate:"required,valSlug"`
	}
)

// Cart Service

type (
	CartInsertRequest struct {
		IdentifierID    uuid.UUID      `json:"identifier_id" validate:"required"`
		ProductID       uuid.UUID      `json:"product_id" validate:"required"`
		ProductQuantity uint16         `json:"product_quantity" validate:"required"`
		Metadata        map[string]any `json:"metadata" validate:"-"`
	}

	CartListRequest struct {
		Token  string `json:"token" validate:"-"`
		Limit  uint16 `json:"limit" validate:"required"`
		Search string `json:"search" validate:"-"`
	}

	CartPatchRequest struct {
		CartID          uuid.UUID `json:"cart_id" validate:"required"`
		IdentifierID    uuid.UUID `json:"identifier_id" validate:"required"`
		ProductID       uuid.UUID `json:"product_id" validate:"required"`
		ProductQuantity uint16    `json:"product_quantity" validate:"required"`
	}

	CartDeleteRequest struct {
		CartID       uuid.UUID `json:"cart_id" validate:"required"`
		IdentifierID uuid.UUID `json:"identifier_id" validate:"required"`
	}
)

// Order Service

type (
	OrderCheckoutRequest struct {
		IdentifierID uuid.UUID  `json:"identifier_id" validate:"required"`
		CartIDs      uuid.UUIDs `json:"cart_ids" validate:"required"`
		FullName     string     `json:"full_name" validate:"required"`
		Email        string     `json:"email" validate:"required,email"`
		PhoneNumber  string     `json:"phone_number" validate:"valPhoneNumber"`
	}

	OrderCheckRequest struct {
		OrderID uuid.UUID `json:"order_id" validate:"required"`
		Email   string    `json:"email" validate:"required,email"`
	}
)
