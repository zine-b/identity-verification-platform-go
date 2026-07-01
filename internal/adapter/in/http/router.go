package httpin

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/postgres"
	"github.com/zineb-b/identity-verification-platform-go/internal/application/service"
)
func NewRouter(db *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()

	
	healthHandler := NewHealthHandler(db)

	userRepository := postgres.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	
	//authUseCase := portin.N
	authHandler := NewAuthHandler(authService)

	mux.HandleFunc("GET /health", healthHandler.Health)
	mux.HandleFunc("POST /auth/signup", authHandler.Signup)

	return Chain(
		mux,
		RecoveryMiddleware,
		RequestIDMiddleware,
		LoggingMiddleware,
	)
}