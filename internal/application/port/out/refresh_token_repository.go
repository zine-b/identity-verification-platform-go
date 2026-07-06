package out

import (
	"context"

	"github.com/zineb-b/identity-verification-platform-go/internal/domain"
)

type RefreshTokenRepository interface {
	Save(ctx context.Context, token *domain.RefreshToken) error
	FindByHash(ctx context.Context, tokenHash string) (*domain.RefreshToken, error)
	Revoke(ctx context.Context, tokenID string) error
}
