package httpin

import (
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
)
func NewRouter(db *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()

	healthHandler := NewHealthHandler(db)

	mux.HandleFunc("GET /health", healthHandler.Health)

	return Chain(
		mux,
		RecoveryMiddleware,
		RequestIDMiddleware,
		LoggingMiddleware,
	)
}