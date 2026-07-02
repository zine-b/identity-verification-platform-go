package httpin

import (
	"net/http"
)

//to avoid cyclic depandancy with container
type Handlers struct {
	HealthHandler *HealthHandler
	AuthHandler   *AuthHandler
}

func NewRouter(handlers Handlers) http.Handler {
	mux := http.NewServeMux()

	healthHandler := handlers.HealthHandler
	authHandler := handlers.AuthHandler

	mux.HandleFunc("GET /health", healthHandler.Health)
	mux.HandleFunc("POST /auth/signup", authHandler.Signup)
	mux.HandleFunc("POST /auth/login", authHandler.Login)

	return Chain(
		mux,
		RecoveryMiddleware,
		RequestIDMiddleware,
		LoggingMiddleware,
	)
}
