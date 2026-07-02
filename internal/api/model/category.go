package model

import (
	"time"

	"github.com/wigata-intech/logres/internal/shared/slugify"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Slug      string    `json:"slug" db:"slug"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewCategory(
	name string,
) (*Category, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Category{
		ID:        id,
		Slug:      slugify.Slugify(name), // TODO: Add Validation Slugify when insert failed
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
