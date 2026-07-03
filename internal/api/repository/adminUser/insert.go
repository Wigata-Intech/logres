package adminUser

import (
	"context"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/model"
)

func (r *adminUserRepository) Create(ctx context.Context, adminUser *model.AdminUser) error {
	_, err := r.db.ExecContext(ctx, insertAdminUser,
		adminUser.ID,
		adminUser.FullName,
		adminUser.Email,
		adminUser.Password,
		adminUser.Status,
		adminUser.LastLoginAt,
		adminUser.CreatedAt,
		adminUser.UpdatedAt,
		adminUser.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("adminUserRepository.Create: %w", err)
	}
	return nil
}
