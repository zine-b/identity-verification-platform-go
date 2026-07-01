package in

import "context"

type SignupCommand struct {
	Email    string
	Password string
}

type SignupResult struct {
	UserID string
	Email  string
	Status string
}

type AuthUseCase interface {
	Signup(ctx context.Context, cmd SignupCommand) (*SignupResult, error)
}