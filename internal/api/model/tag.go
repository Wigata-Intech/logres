package model

import "time"

type Tag struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"` // TODO: Add tag name validation
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewTag(
	name string,
) (*Tag, error) {
	now := time.Now()

	// TODO: Add slug generation

	return &Tag{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
