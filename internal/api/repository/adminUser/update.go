package adminUser

import (
	"context"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

func (r *adminUserRepository) Update(ctx context.Context, adminUser *model.AdminUser) error {
	result, err := r.db.ExecContext(ctx, updateAdminUser,
		adminUser.FullName,
		adminUser.Email,
		adminUser.Password,
		adminUser.Status,
		adminUser.LastLoginAt,
		adminUser.UpdatedAt,
		adminUser.DeletedAt,
		adminUser.ID,
	)
	if err != nil {
		return fmt.Errorf("adminUserRepository.Update: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("adminUserRepository.Update: rows affected: %w", err)
	}
	if affected == 0 {
		return dbx.ErrNotFound
	}
	return nil
}
