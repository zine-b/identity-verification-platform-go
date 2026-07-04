package apperror

import (
	"errors"
)

var (
	ErrEmailRequired        = errors.New("email is required")
	ErrPasswordRequired     = errors.New("password is required")
	ErrPasswordTooShort     = errors.New("password must be at least 8 characters")
	ErrFailedToHashPassword = errors.New("failed to hash password")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordTooWeak      = errors.New("password must contain uppercase, lowercase and digit")
	ErrInvalidCredentials   = errors.New("invalid credentials")
)
