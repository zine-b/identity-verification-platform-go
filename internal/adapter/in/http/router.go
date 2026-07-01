package httpin

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	healthHandler := NewHealthHandler()

	mux.HandleFunc("GET /health", healthHandler.Health)

	return Chain(
		mux,
		RecoveryMiddleware,
		RequestIDMiddleware,
		LoggingMiddleware,
	)
}