package model

import (
	"time"

	"github.com/google/uuid"
)

type AdminPasswordReset struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	AdminUserID  uuid.UUID  `json:"admin_user_id" db:"admin_user_id"`
	OTPHash      string     `json:"-" db:"otp_hash"`
	Purpose      string     `json:"purpose" db:"purpose"`
	AttemptCount int        `json:"attempt_count" db:"attempt_count"`
	ExpiresAt    time.Time  `json:"expires_at" db:"expires_at"`
	UsedAt       *time.Time `json:"used_at,omitempty" db:"used_at"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
}

func NewAdminPasswordReset(adminUserID uuid.UUID, otpHash, purpose string, ttl time.Duration) (*AdminPasswordReset, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &AdminPasswordReset{
		ID:          id,
		AdminUserID: adminUserID,
		OTPHash:     otpHash,
		Purpose:     purpose,
		ExpiresAt:   now.Add(ttl),
		CreatedAt:   now,
	}, nil
}
