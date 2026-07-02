package adminUser

import (
	"context"

	"github.com/wigata-intech/logres/internal/api/model"
)

func (r *adminUserRepository) GetByEmail(ctx context.Context, email string) (*model.AdminUser, error) {
	panic("not implemented") // TODO: Implement
}
