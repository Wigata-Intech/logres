package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AdminRefreshToken struct {
	ID          uuid.UUID                 `json:"id" db:"id"`
	AdminUserID uuid.UUID                 `json:"admin_user_id" db:"admin_user_id"`
	TokenHash   string                    `json:"-" db:"token_hash"`
	FamilyID    uuid.UUID                 `json:"family_id" db:"family_id"`
	ReplacedBy  *uuid.UUID                `json:"replaced_by,omitempty" db:"replaced_by"`
	ExpiresAt   time.Time                 `json:"expires_at" db:"expires_at"`
	RevokedAt   *time.Time                `json:"revoked_at,omitempty" db:"revoked_at"`
	CreatedAt   time.Time                 `json:"created_at" db:"created_at"`
	Metadata    AdminRefreshTokenMetadata `json:"-" db:"metadata"` // PII (contains device.ip_address) — never serialized
	LastUsedAt  time.Time                 `json:"last_used_at" db:"last_used_at"`
}

// AdminRefreshTokenMetadata is the JSON envelope stored per token. Device is
// the only key today; future per-token data (location/login/risk) lands as
// sibling keys with no schema change. It binds and scans as a single JSON
// column.
type AdminRefreshTokenMetadata struct {
	Device AdminRefreshTokenDevice `json:"device"`
}

// AdminRefreshTokenDevice groups the client metadata captured at token
// creation, used for the future active-devices list and per-device revoke.
type AdminRefreshTokenDevice struct {
	UserAgent   string  `json:"user_agent,omitempty"`
	IPAddress   string  `json:"ip_address,omitempty"`
	DeviceLabel *string `json:"device_label,omitempty"`
}

func (m AdminRefreshTokenMetadata) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("AdminRefreshTokenMetadata.Value: %w", err)
	}
	return string(b), nil
}

func (m *AdminRefreshTokenMetadata) Scan(src any) error {
	if src == nil {
		return nil
	}

	var raw string
	switch v := src.(type) {
	case string:
		raw = v
	case []byte:
		raw = string(v)
	default:
		return fmt.Errorf("AdminRefreshTokenMetadata.Scan: unsupported type %T", src)
	}

	if raw == "" {
		*m = AdminRefreshTokenMetadata{}
		return nil
	}

	return json.Unmarshal([]byte(raw), m)
}

// NewAdminRefreshToken starts a new rotation family. Pass an existing
// familyID when rotating a token so the successor stays linked to the
// original login.
func NewAdminRefreshToken(adminUserID uuid.UUID, tokenHash string, familyID uuid.UUID, ttl time.Duration, device AdminRefreshTokenDevice) (*AdminRefreshToken, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &AdminRefreshToken{
		ID:          id,
		AdminUserID: adminUserID,
		TokenHash:   tokenHash,
		FamilyID:    familyID,
		ExpiresAt:   now.Add(ttl),
		CreatedAt:   now,
		Metadata:    AdminRefreshTokenMetadata{Device: device},
		LastUsedAt:  now,
	}, nil
}
