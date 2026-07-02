package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	portout "github.com/zineb-b/identity-verification-platform-go/internal/application/port/out"
)

type JWTManager struct {
	// slice de bytes
	secret []byte
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{
		secret: []byte(secret),
	}
}

func (m *JWTManager) GenerateAccessToken(claims portout.TokenClaims, ttl time.Duration) (string, error) {
	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    claims.UserID,
		"email":  claims.Email,
		"status": claims.Status,
		"iat":    now.Unix(),
		"exp":    now.Add(ttl).Unix(),
	})

	return token.SignedString(m.secret)
}