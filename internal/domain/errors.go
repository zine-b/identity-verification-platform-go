package domain

import (
	"errors"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserAlreadyExists = errors.New("user already exists")
)
