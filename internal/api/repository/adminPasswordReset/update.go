package adminPasswordReset

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/wigata-intech/logres/internal/shared/dbx"
)

func (r *adminPasswordResetRepository) IncrementAttempt(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, incrementAdminPasswordResetAttempt, id)
	if err != nil {
		return fmt.Errorf("adminPasswordResetRepository.IncrementAttempt: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("adminPasswordResetRepository.IncrementAttempt: rows affected: %w", err)
	}
	if affected == 0 {
		return dbx.ErrNotFound
	}
	return nil
}

func (r *adminPasswordResetRepository) MarkUsed(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, markAdminPasswordResetUsed, time.Now(), id)
	if err != nil {
		return fmt.Errorf("adminPasswordResetRepository.MarkUsed: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("adminPasswordResetRepository.MarkUsed: rows affected: %w", err)
	}
	if affected == 0 {
		return dbx.ErrNotFound
	}
	return nil
}

func (r *adminPasswordResetRepository) InvalidateAllForUser(ctx context.Context, adminUserID uuid.UUID, purpose string) error {
	_, err := r.db.ExecContext(ctx, invalidateAdminPasswordResetsForUser, time.Now(), adminUserID, purpose)
	if err != nil {
		return fmt.Errorf("adminPasswordResetRepository.InvalidateAllForUser: %w", err)
	}
	return nil
}
