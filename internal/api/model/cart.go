package model

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID           uuid.UUID      `json:"id" db:"id"`
	IdentifierID uuid.UUID      `json:"identifier_id" db:"identifier_id"` // ID for identifier / group. For now, will use device_id but in the future use user_id or session_id
	ProductID    uuid.UUID      `json:"product_id" db:"product_id"`
	Quantity     uint16         `json:"quantity" db:"quantity"`
	Order        uint8          `json:"order" db:"order"`
	Metadata     map[string]any `json:"metadata,omitempty" db:"metadata"` // Additional metadata for the cart item, such as customizations or options
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

func NewCart(
	identifierID, productID uuid.UUID,
	quantity uint16,
	order uint8,
) (*Cart, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Cart{
		ID:           id,
		IdentifierID: identifierID,
		ProductID:    productID,
		Quantity:     quantity,
		Order:        order,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

func (c *Cart) SetMetadata(key string, value any) {
	_, ok := c.Metadata[key]
	if !ok {
		c.Metadata = make(map[string]any)
	}

	c.Metadata[key] = value
	c.UpdatedAt = time.Now()
}

func (c *Cart) GetMetadata(key string) any {
	return c.Metadata[key]
}
