package domain

import (
	"time"
	"github.com/zineb-b/identity-verification-platform-go/internal/application/apperror"

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

func NewUser(id string, email string, passwordHash string, now time.Time) (*User, error) {
	if email == "" {
		return nil, apperror.ErrEmailRequired
	}

	if passwordHash == "" {
		return nil, apperror.ErrPasswordRequired
	}

	return &User{
		ID:           id,
		Email:        email,
		PasswordHash: passwordHash,
		Status:       UserStatusPending,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil

}
