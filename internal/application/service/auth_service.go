package service

import (
	"context"
	"strings"

	portin "github.com/zineb-b/identity-verification-platform-go/internal/application/port/in"
	portout "github.com/zineb-b/identity-verification-platform-go/internal/application/port/out"
	"github.com/zineb-b/identity-verification-platform-go/internal/domain"
)

type AuthService struct {
	userRepo portout.UserRepository
	hasher portout.PasswordHasher
	idGenerator portout.IDGenerator
	clock portout.Clock
}

	func NewAuthService(
		userRepo portout.UserRepository, 
		hasher portout.PasswordHasher, 
		iDGenerator portout.IDGenerator, 
		clock portout.Clock,
	) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		hasher: hasher,
		idGenerator: iDGenerator,
		clock: clock,
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

	existingUser, err := s.userRepo.FindByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	// hash password
	passwordHash, err := s.hasher.Hash(cmd.Password)
	if err != nil {
		return nil, ErrFailedToHashPassword
	}
	user, err := domain.NewUser(
		s.idGenerator.NewID(),
		email,
		string(passwordHash),
		s.clock.Now(),

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
