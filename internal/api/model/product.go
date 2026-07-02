package model

import (
	"time"

	"github.com/wigata-intech/logres/internal/api/constant"
	"github.com/wigata-intech/logres/internal/shared/slugify"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID          uuid.UUID              `json:"id" db:"id"`
	Slug        string                 `json:"slug" db:"slug"`
	Name        string                 `json:"name" db:"name"`
	Description string                 `json:"description" db:"description"`
	Price       decimal.Decimal        `json:"price" db:"price"`
	Status      constant.ProductStatus `json:"status" db:"status"`
	CreatedAt   time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time             `json:"deleted_at,omitempty" db:"deleted_at"`
}

type ProductFile struct {
	ProductID   uuid.UUID `json:"product_id" db:"product_id"`
	FileID      uuid.UUID `json:"file_id" db:"file_id"`
	Order       uint8     `json:"order" db:"order"`
	IsThumbnail bool      `json:"is_thumbnail" db:"is_thumbnail"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ProductCategory struct {
	ProductID  uuid.UUID `json:"product_id" db:"product_id"`
	CategoryID uuid.UUID `json:"category_id" db:"category_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type ProductTag struct {
	ProductID uuid.UUID `json:"product_id" db:"product_id"`
	TagID     uuid.UUID `json:"tag_id" db:"tag_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewProduct(
	name, description string,
	price decimal.Decimal,
) (*Product, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Product{
		ID:          id,
		Slug:        slugify.Slugify(name), // TODO: Add Validation Slugify when insert failed
		Name:        name,
		Description: description,
		Price:       price,
		Status:      constant.ProductStatusDraft,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (p *Product) Publish() {
	p.Status = constant.ProductStatusPublished
	p.UpdatedAt = time.Now()
}

func (p *Product) Archive() {
	p.Status = constant.ProductStatusArchived
	p.UpdatedAt = time.Now()
}

func NewProductFile(
	productID uuid.UUID,
	fileID uuid.UUID,
	order uint8,
	isThumbnail bool,
) (*ProductFile, error) {
	now := time.Now()

	return &ProductFile{
		ProductID:   productID,
		FileID:      fileID,
		Order:       order,
		IsThumbnail: isThumbnail,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func NewProductCategory(
	productID uuid.UUID,
	categoryID uuid.UUID,
) (*ProductCategory, error) {
	now := time.Now()

	return &ProductCategory{
		ProductID:  productID,
		CategoryID: categoryID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}

func NewProductTag(
	productID uuid.UUID,
	tagID uuid.UUID,
) (*ProductTag, error) {
	now := time.Now()

	return &ProductTag{
		ProductID: productID,
		TagID:     tagID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
