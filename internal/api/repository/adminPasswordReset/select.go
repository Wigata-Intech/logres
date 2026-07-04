package adminPasswordReset

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

func (r *adminPasswordResetRepository) GetActiveByUser(ctx context.Context, adminUserID uuid.UUID, purpose string) (*model.AdminPasswordReset, error) {
	var reset model.AdminPasswordReset
	err := r.db.QueryRowContext(ctx, selectActiveAdminPasswordResetByUser, adminUserID, purpose).Scan(
		&reset.ID,
		&reset.AdminUserID,
		&reset.OTPHash,
		&reset.Purpose,
		&reset.AttemptCount,
		&reset.ExpiresAt,
		&reset.UsedAt,
		&reset.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, dbx.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("adminPasswordResetRepository.GetActiveByUser: %w", err)
	}
	return &reset, nil
}
