package adminRefreshToken

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/wigata-intech/logres/internal/api/model"
	"github.com/wigata-intech/logres/internal/shared/dbx"
)

func (r *adminRefreshTokenRepository) GetByHash(ctx context.Context, tokenHash string) (*model.AdminRefreshToken, error) {
	var token model.AdminRefreshToken
	err := r.db.QueryRowContext(ctx, selectAdminRefreshTokenByHash, tokenHash).Scan(
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
		return nil, fmt.Errorf("adminRefreshTokenRepository.GetByHash: %w", err)
	}
	return &token, nil
}
