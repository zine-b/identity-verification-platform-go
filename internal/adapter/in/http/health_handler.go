package httpin

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type HealthHandler struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

type HealthResponse struct {
	Status   string `json:"status"`
	Postgres string `json:"postgres"`
	Redis    string `json:"redis"`
}

func NewHealthHandler(db *pgxpool.Pool, redis *redis.Client) *HealthHandler {
	return &HealthHandler{
		db:    db,
		redis: redis,
	}
}

// http handler
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	postgresStatus := "ok"

	if err := h.db.Ping(r.Context()); err != nil {
		postgresStatus = "down"
	}

	redisStatus := "ok"
	if err := h.redis.Ping(r.Context()).Err(); err != nil {
		redisStatus = "down"
	}

	response := HealthResponse{
		Status:   "ok",
		Postgres: postgresStatus,
		Redis:    redisStatus,
	}

	writeJSON(w, http.StatusOK, response)
}
