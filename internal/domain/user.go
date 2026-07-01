package domain

import (
	"time"
)

type UserStatus string

const (
	UserStatusPending UserStatus = "pending"
	UserStatusActive  UserStatus = "active"
	UserStatusBlocked UserStatus = "blocked"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Status       UserStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(id string, email string, passwordHash string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}

	if passwordHash == "" {
		return nil, ErrInvalidPassword
	}

	now := time.Now()

	return &User{
		ID:           id,
		Email:        email,
		PasswordHash: passwordHash,
		Status:       UserStatusPending,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil

}
