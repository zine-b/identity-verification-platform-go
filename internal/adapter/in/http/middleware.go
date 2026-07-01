package httpin

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"time"
)

type contextKey string

const requestIDKey contextKey = "request_id"

// fonction qui prend handler HTTP et retourne un nouveau handler HTTP. (router mux)
// handler d’origine
//
//	↓
//
// middleware(handler)
//
//	↓
//
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

		// attend juqu'a les autres Middlewares terminent
		log.Printf(
			"request_id=%s method=%s path=%s duration=%s",
			GetRequestID(r.Context()),
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

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = generateRequestID()
		}

		w.Header().Set("X-Request-Id", requestID)

		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return ""
	}

	return requestID
}

func generateRequestID() string {
	bytes := make([]byte, 16)

	if _, err := rand.Read(bytes); err != nil {
		return time.Now().Format("20060102150405")
	}

	return hex.EncodeToString(bytes)
}
