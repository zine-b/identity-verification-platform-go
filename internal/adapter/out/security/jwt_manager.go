package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	portout "github.com/zineb-b/identity-verification-platform-go/internal/application/port/out"
)

type JWTManager struct {
	// slice de bytes
	// la clé de signature du token qui existe dans la conf est en string
	// mais jwt attend un byte 
	// donc object secret c pour convertir string --> byte 
	secret []byte
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{
		secret: []byte(secret),
	}
}

func (m *JWTManager) GenerateAccessToken(claims portout.TokenClaims, ttl time.Duration) (string, error) {
	now := time.Now()

	// generer le token 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    claims.UserID,
		"email":  claims.Email,
		"status": claims.Status,
		"iat":    now.Unix(),
		"exp":    now.Add(ttl).Unix(),
	})

	// signer le token 
	return token.SignedString(m.secret)
}

func (m *JWTManager) ValidateAccessToken(tokenString string) (*portout.TokenClaims, error) {
	// verifie que ce token a été creer avec cette signature
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return m.secret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return &portout.TokenClaims{
		UserID: claims["sub"].(string),
		Email:  claims["email"].(string),
		Status: claims["status"].(string),
	}, nil
}