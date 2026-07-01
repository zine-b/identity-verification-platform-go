package httpin

import (
	"log"
	"net/http"
	"time"
)


// fonction qui prend handler HTTP et retourne un nouveau handler HTTP. (router mux)
// handler d’origine
//        ↓
// middleware(handler)
//        ↓
// handler amélioré
type Middleware func(http.Handler) http.Handler

// appliquer plusieurs middlewares autour d’un handler.
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

// ajoute du logging autour d’une requête.
func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"method=%s path=%s duration=%s",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

// protéger ton serveur contre les panic.
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered: %v", err)
				writeJSON(w, http.StatusInternalServerError, map[string]string{
					"error": "internal server error",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}