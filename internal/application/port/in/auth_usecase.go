package in

import "context"

type SignupCommand struct {
	Email    string
	Password string
}

type SignupResult struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type AuthUseCase interface {
	Signup(ctx context.Context, cmd SignupCommand) (*SignupResult, error)
}