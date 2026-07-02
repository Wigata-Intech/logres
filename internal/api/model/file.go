package model

import (
	"time"

	"github.com/wigata-intech/logres/internal/api/constant"

	"github.com/google/uuid"
)

type File struct {
	ID        uuid.UUID         `json:"id" db:"id"`
	Name      string            `json:"name" db:"name"`
	Path      string            `json:"path" db:"path"`
	MimeType  constant.MimeType `json:"mime_type" db:"mime_type"`
	Blob      *string           `json:"blob,omitempty" db:"blob"` // TODO: Add blob storage
	IsPublic  bool              `json:"is_public" db:"is_public"`
	CreatedAt time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time        `json:"deleted_at,omitempty" db:"deleted_at"`
}

func NewFile(
	name, path string,
	mimeType constant.MimeType,
	blob *string,
	isPublic bool,
) (*File, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &File{
		ID:        id,
		Name:      name,
		Path:      path,
		MimeType:  mimeType,
		Blob:      blob,
		IsPublic:  isPublic,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
