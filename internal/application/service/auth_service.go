package service

import (
	"context"
	"strings"
	"errors"

	portin "github.com/zineb-b/identity-verification-platform-go/internal/application/port/in"
	portout "github.com/zineb-b/identity-verification-platform-go/internal/application/port/out"
	"github.com/zineb-b/identity-verification-platform-go/internal/domain"
	"github.com/zineb-b/identity-verification-platform-go/internal/application/apperror"
	"github.com/zineb-b/identity-verification-platform-go/internal/application/validation"

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
		idGenerator portout.IDGenerator, 
		clock portout.Clock,
	) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		hasher: hasher,
		idGenerator: idGenerator,
		clock: clock,
	}
}

func (s *AuthService) Signup(ctx context.Context, cmd portin.SignupCommand) (*portin.SignupResult, error) {
	email := strings.TrimSpace(strings.ToLower(cmd.Email))

	if !validation.IsValidEmail(email) {
		return nil, apperror.ErrInvalidEmail
	}

	if cmd.Password == "" {
		return nil, apperror.ErrPasswordRequired
	}

	if len(cmd.Password) < 8 {
		return nil, apperror.ErrPasswordTooShort
	}

	existingUser, err := s.userRepo.FindByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, apperror.ErrUserAlreadyExists
	}

	if err != nil && !errors.Is(err, apperror.ErrUserNotFound) {
		return nil, err
	}

	// hash password
	passwordHash, err := s.hasher.Hash(cmd.Password)
	if err != nil {
		return nil, apperror.ErrFailedToHashPassword
	}
	user, err := domain.NewUser(
		s.idGenerator.NewID(),
		email,
		passwordHash,
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
