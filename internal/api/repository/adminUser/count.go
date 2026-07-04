package adminUser

import (
	"context"
	"fmt"
)

func (r *adminUserRepository) CountAll(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, countAdminUsers).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("adminUserRepository.CountAll: %w", err)
	}
	return count, nil
}
