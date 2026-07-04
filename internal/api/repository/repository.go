package repository

import (
	"context"

	"github.com/wigata-intech/logres/internal/api/model"
	sharedModel "github.com/wigata-intech/logres/internal/shared/model"

	"github.com/google/uuid"
)

type IAdminUserRepository interface {
	Create(ctx context.Context, adminUser *model.AdminUser) error
	GetByEmail(ctx context.Context, email string) (*model.AdminUser, error)
	Update(ctx context.Context, adminUser *model.AdminUser) error
	CountAll(ctx context.Context) (int, error)
}

type IAdminPasswordResetRepository interface {
	Create(ctx context.Context, reset *model.AdminPasswordReset) error
	GetActiveByUser(ctx context.Context, adminUserID uuid.UUID, purpose string) (*model.AdminPasswordReset, error)
	IncrementAttempt(ctx context.Context, id uuid.UUID) error
	MarkUsed(ctx context.Context, id uuid.UUID) error
	InvalidateAllForUser(ctx context.Context, adminUserID uuid.UUID, purpose string) error
}

type IAdminRefreshTokenRepository interface {
	Insert(ctx context.Context, token *model.AdminRefreshToken) error
	GetByHash(ctx context.Context, tokenHash string) (*model.AdminRefreshToken, error)
	Rotate(ctx context.Context, oldID, newID uuid.UUID) error
	RevokeFamily(ctx context.Context, familyID uuid.UUID) error
	RevokeAllForUser(ctx context.Context, adminUserID uuid.UUID) error
}

type ICartRepository interface {
	Insert(ctx context.Context, cart *model.Cart) error
	List(ctx context.Context, pagination *sharedModel.Pagination, identifierID uuid.UUID) ([]model.Cart, *sharedModel.Pagination, error)
	GetByIdentifierID(ctx context.Context, identifierID uuid.UUID) (*model.Cart, error)
	Update(ctx context.Context, cart *model.Cart) error
	Delete(ctx context.Context, cart *model.Cart) error
}

type ICategoryRepository interface {
	Create(ctx context.Context, category *model.Category) error
	List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.Category, *sharedModel.Pagination, error)
	GetBySlug(ctx context.Context, slug string) (*model.Category, error)
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, category *model.Category) error
}

type IFileRepository interface {
	Insert(ctx context.Context, file *model.File) error
	List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.File, *sharedModel.Pagination, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.File, error)
	Update(ctx context.Context, file *model.File) error
	Delete(ctx context.Context, file *model.File) error
}

type IOrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.Order, *sharedModel.Pagination, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
	Delete(ctx context.Context, order *model.Order) error
}

type IProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.Product, *sharedModel.Pagination, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Product, error)
	GetBySlug(ctx context.Context, slug string) (*model.Product, error)
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, product *model.Product) error

	// Files
	FileAssign(ctx context.Context, productFile *model.ProductFile) error
	FileUnassign(ctx context.Context, productFile *model.ProductFile) error
	FileListByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductFile, error)

	// Category and Tags
	CategoryAssign(ctx context.Context, productCategory *model.ProductCategory) error
	CategoryUnassign(ctx context.Context, productCategory *model.ProductCategory) error
	CategoryListByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductCategory, error)
	TagAssign(ctx context.Context, productTag *model.ProductTag) error
	TagUnassign(ctx context.Context, productTag *model.ProductTag) error
	TagListByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductTag, error)
}

type ITagRepository interface {
	Create(ctx context.Context, tag *model.Tag) error
	List(ctx context.Context, pagination *sharedModel.Pagination, search string) ([]model.Tag, *sharedModel.Pagination, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Tag, error)
	Update(ctx context.Context, tag *model.Tag) error
	Delete(ctx context.Context, tag *model.Tag) error
}
