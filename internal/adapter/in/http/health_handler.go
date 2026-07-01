package httpin

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type HealthHandler struct {
	db *pgxpool.Pool
}

type HealthResponse struct {
	Status   string `json:"status"`
	Postgres string `json:"postgres"`
}

func NewHealthHandler(db *pgxpool.Pool) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

// http handler
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	postgresStatus := "ok"

	if err := h.db.Ping(r.Context()); err != nil {
		postgresStatus = "down"
	}

	response := HealthResponse{
		Status:   "ok",
		Postgres: postgresStatus,
	}

	writeJSON(w, http.StatusOK, response)
}
