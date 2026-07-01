package out

import (
	"context"
	"github.com/zineb-b/identity-verification-platform-go/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}