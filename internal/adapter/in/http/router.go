package httpin

import (
	portout "github.com/zineb-b/identity-verification-platform-go/internal/application/port/out"
	"log/slog"
	"net/http"
)

// to avoid cyclic depandancy with container
type Handlers struct {
	HealthHandler *HealthHandler
	AuthHandler   *AuthHandler
	TokenManager  portout.TokenManager
	Logger        *slog.Logger
}

func NewRouter(handlers Handlers) http.Handler {
	mux := http.NewServeMux()

	healthHandler := handlers.HealthHandler
	authHandler := handlers.AuthHandler
	meHandler := AuthMiddleware(handlers.TokenManager)(
		http.HandlerFunc(handlers.AuthHandler.Me),
	)

	mux.Handle("GET /me", meHandler)
	mux.HandleFunc("GET /health", healthHandler.Health)
	mux.HandleFunc("POST /auth/signup", authHandler.Signup)
	mux.HandleFunc("POST /auth/login", authHandler.Login)

	return Chain(
		mux,
		RecoveryMiddleware,
		RequestIDMiddleware,
		LoggingMiddleware(handlers.Logger),
	)
}
