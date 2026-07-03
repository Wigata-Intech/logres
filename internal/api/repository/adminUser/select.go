package adminUser

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

func (r *adminUserRepository) GetByEmail(ctx context.Context, email string) (*model.AdminUser, error) {
	var adminUser model.AdminUser
	err := r.db.QueryRowContext(ctx, selectAdminUserByEmail, email).Scan(
		&adminUser.ID,
		&adminUser.FullName,
		&adminUser.Email,
		&adminUser.Password,
		&adminUser.Status,
		&adminUser.LastLoginAt,
		&adminUser.CreatedAt,
		&adminUser.UpdatedAt,
		&adminUser.DeletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, dbx.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("adminUserRepository.GetByEmail: %w", err)
	}
	return &adminUser, nil
}
