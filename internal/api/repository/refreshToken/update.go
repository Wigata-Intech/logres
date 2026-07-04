package refreshToken

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
func (r *refreshTokenRepository) Rotate(ctx context.Context, oldID, newID uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, rotateRefreshToken, newID, time.Now(), oldID)
	if err != nil {
		return fmt.Errorf("refreshTokenRepository.Rotate: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("refreshTokenRepository.Rotate: rows affected: %w", err)
	}
	if affected == 0 {
		return dbx.ErrNotFound
	}
	return nil
}

func (r *refreshTokenRepository) RevokeFamily(ctx context.Context, familyID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, revokeRefreshTokenFamily, time.Now(), familyID)
	if err != nil {
		return fmt.Errorf("refreshTokenRepository.RevokeFamily: %w", err)
	}
	return nil
}

func (r *refreshTokenRepository) RevokeAllForUser(ctx context.Context, adminUserID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, revokeAllRefreshTokensForUser, time.Now(), adminUserID)
	if err != nil {
		return fmt.Errorf("refreshTokenRepository.RevokeAllForUser: %w", err)
	}
	return nil
}
