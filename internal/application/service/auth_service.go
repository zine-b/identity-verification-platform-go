package service

import (
	"context"

	portin "github.com/zineb-b/identity-verification-platform-go/internal/application/port/in"
	portout "github.com/zineb-b/identity-verification-platform-go/internal/application/port/out"
)


type AuthService struct {
	userRepo portout.UserRepository
}

func NewAuthService(userRepo portout.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Signup(ctx context.Context, cmd portin.SignupCommand)(*portin.SignupResult, error){
	return &portin.SignupResult{
		UserID: "todo",
		Email:  cmd.Email,
		Status: "pending",
	}, nil
}