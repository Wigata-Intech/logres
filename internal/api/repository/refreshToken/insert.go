package refreshToken

import (
	"context"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/model"
)

func (r *refreshTokenRepository) Insert(ctx context.Context, token *model.RefreshToken) error {
	_, err := r.db.ExecContext(ctx, insertRefreshToken,
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
		return fmt.Errorf("refreshTokenRepository.Insert: %w", err)
	}
	return nil
}
