package adminPasswordReset

import (
	"context"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/model"
)

func (r *adminPasswordResetRepository) Create(ctx context.Context, reset *model.AdminPasswordReset) error {
	_, err := r.db.ExecContext(ctx, insertAdminPasswordReset,
		reset.ID,
		reset.AdminUserID,
		reset.OTPHash,
		reset.Purpose,
		reset.AttemptCount,
		reset.ExpiresAt,
		reset.UsedAt,
		reset.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("adminPasswordResetRepository.Create: %w", err)
	}
	return nil
}
