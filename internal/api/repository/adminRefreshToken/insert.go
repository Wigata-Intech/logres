package adminRefreshToken

import (
	"context"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/model"
)

func (r *adminRefreshTokenRepository) Insert(ctx context.Context, token *model.AdminRefreshToken) error {
	_, err := r.db.ExecContext(ctx, insertAdminRefreshToken,
		token.ID,
		token.AdminUserID,
		token.TokenHash,
		token.FamilyID,
		token.ReplacedBy,
		token.ExpiresAt,
		token.RevokedAt,
		token.CreatedAt,
		token.Metadata,
		token.LastUsedAt,
	)
	if err != nil {
		return fmt.Errorf("adminRefreshTokenRepository.Insert: %w", err)
	}
	return nil
}
