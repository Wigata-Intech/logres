package model

import (
	"time"

	"github.com/wigata-intech/logres/internal/api/constant"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Admin User (Auth Service)

type (
	AdminSetupResponse struct {
		Message string `json:"message"`
	}

	// AdminTokenResponse is the RFC 6749 §5.1 access-token payload.
	AdminTokenResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}

	AdminForgotPasswordResponse struct {
		Message string `json:"message"`
	}

	AdminChangePasswordResponse struct {
		Message string `json:"message"`
	}

	AdminResetPasswordResponse struct {
		Message string `json:"message"`
	}
)

// Admin File Management Service

type (
	AdminFileResponse struct {
		ID        uuid.UUID         `json:"id"`
		Name      string            `json:"name"`
		Path      string            `json:"path"`
		MimeType  constant.MimeType `json:"mime_type"`
		Blob      *string           `json:"blob,omitempty"`
		IsPublic  bool              `json:"is_public"`
		CreatedAt time.Time         `json:"created_at"`
		UpdatedAt time.Time         `json:"updated_at"`
	}

	AdminFileListResponse struct {
		Files     []AdminFileResponse `json:"files"`
		NextToken *string             `json:"next_token"`
		Limit     uint16              `json:"limit"`
		// TODO: Later on do A/B testing do we need total_data / total_page per limit
	}
)

// Admin Product Service

type (
	AdminProductResponse struct {
		ID          uuid.UUID              `json:"id"`
		Slug        string                 `json:"slug"`
		Name        string                 `json:"name"`
		Description string                 `json:"description"`
		Price       decimal.Decimal        `json:"price"`
		Status      constant.ProductStatus `json:"status"`
		CreatedAt   time.Time              `json:"created_at"`
		UpdatedAt   time.Time              `json:"updated_at"`

		Files      []AdminFileResponse     `json:"files,omitempty"`
		Categories []AdminCategoryResponse `json:"categories,omitempty"`
		Tags       []AdminTagResponse      `json:"tags,omitempty"`
	}

	AdminProductListResponse struct {
		Products  []AdminProductResponse `json:"products"`
		NextToken *string                `json:"next_token"`
		Limit     uint16                 `json:"limit"`
		// TODO: Later on do A/B testing do we need total_data / total_page per limit
	}

	// Category and Tags
	AdminCategoryResponse struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	}

	AdminCategoryListResponse struct {
		Categories []AdminCategoryResponse `json:"categories"`
		NextToken  *string                 `json:"next_token"`
		Limit      uint16                  `json:"limit"`
		// TODO: Later on do A/B testing do we need total_data / total_page per limit
	}

	AdminTagResponse struct {
		Name string `json:"name"`
	}

	AdminTagListResponse struct {
		Tags      []AdminTagResponse `json:"tags"`
		NextToken *string            `json:"next_token"`
		Limit     uint16             `json:"limit"`
		// TODO: Later on do A/B testing do we need total_data / total_page per limit
	}
)

// Admin Order Service

type (
	AdminOrderItemResponse struct {
		OrderID    uuid.UUID       `json:"order_id"`
		ProductID  uuid.UUID       `json:"product_id"`
		Quantity   uint16          `json:"quantity"`
		Price      decimal.Decimal `json:"price"`
		TotalPrice decimal.Decimal `json:"total_price"`
		Order      uint8           `json:"order"`
	}

	AdminOrderResponse struct {
		ID            uuid.UUID            `json:"id"`
		DocumentID    uuid.UUID            `json:"document_id"`
		PaymentID     uuid.NullUUID        `json:"payment_id"`
		FullName      string               `json:"full_name"`
		Email         string               `json:"email"`
		PhoneNumber   string               `json:"phone_number"`
		TotalPrice    decimal.Decimal      `json:"total_price"`
		TotalDiscount decimal.Decimal      `json:"total_discount"`
		TotalTax      decimal.Decimal      `json:"total_tax"`
		BilledAmount  decimal.Decimal      `json:"billed_amount"`
		Status        constant.OrderStatus `json:"status"`
		CreatedAt     time.Time            `json:"created_at"`
		UpdatedAt     time.Time            `json:"updated_at"`

		Items []AdminOrderItemResponse `json:"items,omitempty"`
	}

	AdminOrderListResponse struct {
		Orders    []AdminOrderResponse `json:"orders"`
		NextToken *string              `json:"next_token"`
		Limit     uint16               `json:"limit"`
		// TODO: Later on do A/B testing do we need total_data / total_page per limit
	}

	// Cart Service
	AdminCartResponse struct {
		ID           uuid.UUID      `json:"id"`
		IdentifierID uuid.UUID      `json:"identifier_id"`
		ProductID    uuid.UUID      `json:"product_id"`
		Quantity     uint16         `json:"quantity"`
		Order        uint8          `json:"order"`
		Metadata     map[string]any `json:"metadata,omitempty"`
		CreatedAt    time.Time      `json:"created_at"`
		UpdatedAt    time.Time      `json:"updated_at"`
	}

	AdminCartListResponse struct {
		Carts     []AdminCartResponse `json:"carts"`
		NextToken *string             `json:"next_token"`
		Limit     uint16              `json:"limit"`
		// TODO: Later on do A/B testing do we need total_data / total_page per limit
	}
)

// Product Service

type (
	FileResponse struct {
		ID        uuid.UUID         `json:"id"`
		Name      string            `json:"name"`
		Path      string            `json:"path"`
		MimeType  constant.MimeType `json:"mime_type"`
		Blob      *string           `json:"blob,omitempty"`
		CreatedAt time.Time         `json:"created_at"`
		UpdatedAt time.Time         `json:"updated_at"`
	}

	ProductResponse struct {
		ID          uuid.UUID       `json:"id"`
		Slug        string          `json:"slug"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Price       decimal.Decimal `json:"price"`
		CreatedAt   time.Time       `json:"created_at"`
		UpdatedAt   time.Time       `json:"updated_at"`

		Files      []FileResponse     `json:"files,omitempty"`
		Categories []CategoryResponse `json:"categories,omitempty"`
		Tags       []TagRespose       `json:"tags,omitempty"`
	}

	ProductListResponse struct {
		Products  []ProductResponse
		NextToken *string `json:"next_token"`
		Limit     uint16  `json:"limit"`
		// TODO: Later on do A/B testing do we need total_data / total_page per limit
	}

	CategoryResponse struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	}

	CategoryListRespones struct {
		Categories []CategoryResponse `json:"categories"`
		NextToken  *string            `json:"next_token"`
		Limit      uint16             `json:"limit"`
		// TODO: Later on do A/B testing do we need total_data / total_page per limit
	}

	TagRespose struct {
		Name string `json:"name"`
	}
)

// Cart Service

type (
	CartResponse struct {
		ID           uuid.UUID `json:"id"`
		IdentifierID uuid.UUID `json:"identifier_id"`
		ProductID    uuid.UUID `json:"product_id"`
		Quantity     uint16    `json:"quantity"`
		Order        uint8     `json:"order"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	CartListResponse struct {
		Carts     []CartResponse
		NextToken *string `json:"next_token"`
		Limit     uint16  `json:"limit"`
		// TODO: Later on do A/B testing do we need total_data / total_page per limit
	}
)

// Order Service

type (
	// TODO: Need to think about the flow first
	// because after checkout either digital payment
	// or manual payment confirmations
	OrderCheckoutResponse struct {
	}

	OrderResponse struct {
		ID            uuid.UUID            `json:"id"`
		DocumentID    uuid.UUID            `json:"document_id"`
		PaymentID     uuid.NullUUID        `json:"payment_id"`
		FullName      string               `json:"full_name"`
		Email         string               `json:"email"`
		PhoneNumber   string               `json:"phone_number"`
		TotalPrice    decimal.Decimal      `json:"total_price"`
		TotalDiscount decimal.Decimal      `json:"total_discount"`
		TotalTax      decimal.Decimal      `json:"total_tax"`
		BilledAmount  decimal.Decimal      `json:"billed_amount"`
		Status        constant.OrderStatus `json:"status"`
		CreatedAt     time.Time            `json:"created_at"`
		UpdatedAt     time.Time            `json:"updated_at"`

		Items []AdminOrderItemResponse `json:"items,omitempty"`
	}
)
