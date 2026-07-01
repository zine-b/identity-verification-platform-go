package service

import (
	"errors"
)

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordTooShort = errors.New("password must be at least 8 characters")
)
