package model

import (
	"time"

	"github.com/wigata-intech/logres/internal/api/constant"

	"github.com/google/uuid"
)

type AdminUser struct {
	ID          uuid.UUID                `json:"id" db:"id"`
	FullName    string                   `json:"full_name" db:"full_name"`
	Email       string                   `json:"email" db:"email"`
	Password    string                   `json:"password" db:"password"`
	Status      constant.AdminUserStatus `json:"status" db:"status"`
	LastLoginAt *time.Time               `json:"last_login_at,omitempty" db:"last_login_at"`
	CreatedAt   time.Time                `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time               `json:"deleted_at,omitempty" db:"deleted_at"`
}

// IsActive reports whether the account may authenticate. Fetch paths already
// exclude soft-deleted rows, so status is the only remaining gate.
func (u *AdminUser) IsActive() bool {
	return u.Status == constant.AdminUserStatusActive
}

func NewAdminUser(fullName, email string) (*AdminUser, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &AdminUser{
		ID:        id,
		FullName:  fullName,
		Email:     email,
		Status:    constant.AdminUserStatusInactive,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
