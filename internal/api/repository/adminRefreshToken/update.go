package adminRefreshToken

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/wigata-intech/logres/internal/shared/dbx"
)

// Rotate retires oldID in favour of newID: sets replaced_by and revoked_at on
// the old row. It is a no-op if the row is already revoked (reuse-detection
// relies on that distinction, not on this call).
func (r *adminRefreshTokenRepository) Rotate(ctx context.Context, oldID, newID uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, rotateAdminRefreshToken, newID, time.Now(), oldID)
	if err != nil {
		return fmt.Errorf("adminRefreshTokenRepository.Rotate: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("adminRefreshTokenRepository.Rotate: rows affected: %w", err)
	}
	if affected == 0 {
		return dbx.ErrNotFound
	}
	return nil
}

func (r *adminRefreshTokenRepository) RevokeFamily(ctx context.Context, familyID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, revokeAdminRefreshTokenFamily, time.Now(), familyID)
	if err != nil {
		return fmt.Errorf("adminRefreshTokenRepository.RevokeFamily: %w", err)
	}
	return nil
}

func (r *adminRefreshTokenRepository) RevokeAllForUser(ctx context.Context, adminUserID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, revokeAllAdminRefreshTokensForUser, time.Now(), adminUserID)
	if err != nil {
		return fmt.Errorf("adminRefreshTokenRepository.RevokeAllForUser: %w", err)
	}
	return nil
}
