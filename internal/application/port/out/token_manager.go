package out

import "time"

type TokenClaims struct {
	UserID string
	Email  string
	Status string
}

type TokenManager interface {
	GenerateAccessToken(claims TokenClaims, ttl time.Duration) (string, error)
	ValidateAccessToken(token string) (*TokenClaims, error)
}
