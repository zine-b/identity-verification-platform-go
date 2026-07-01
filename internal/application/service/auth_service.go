package service

import (
	"context"
	"strings"

	"github.com/google/uuid"
	portin "github.com/zineb-b/identity-verification-platform-go/internal/application/port/in"
	portout "github.com/zineb-b/identity-verification-platform-go/internal/application/port/out"
	"github.com/zineb-b/identity-verification-platform-go/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo portout.UserRepository
}

func NewAuthService(userRepo portout.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Signup(ctx context.Context, cmd portin.SignupCommand) (*portin.SignupResult, error) {
	email := strings.TrimSpace(strings.ToLower(cmd.Email))

	if email == "" {
		return nil, ErrEmailRequired
	}

	if cmd.Password == "" {
		return nil, ErrPasswordRequired
	}

	if len(cmd.Password) < 8 {
		return nil, ErrPasswordTooShort
	}

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(
		uuid.NewString(),
		email,
		string(passwordHash),
	)
	if err != nil {
		return nil, err
	}

	if err := s.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}

	return &portin.SignupResult{
		UserID: user.ID,
		Email:  user.Email,
		Status: string(user.Status),
	}, nil
}
