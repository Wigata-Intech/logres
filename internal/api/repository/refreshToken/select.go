package refreshToken

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

func (r *refreshTokenRepository) GetByHash(ctx context.Context, tokenHash string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	err := r.db.QueryRowContext(ctx, selectRefreshTokenByHash, tokenHash).Scan(
		&token.ID,
		&token.AdminUserID,
		&token.TokenHash,
		&token.FamilyID,
		&token.ReplacedBy,
		&token.ExpiresAt,
		&token.RevokedAt,
		&token.CreatedAt,
		&token.Metadata,
		&token.LastUsedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, dbx.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("refreshTokenRepository.GetByHash: %w", err)
	}
	return &token, nil
}
