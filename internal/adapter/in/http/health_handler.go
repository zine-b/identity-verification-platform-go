package httpin

import (
	"net/http"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
)
type HealthHandler struct{
	db *pgxpool.Pool
}

type HealthResponse struct{
	Status   string `json:"status"`
	Postgres string `json:"postgres"`
}

func NewHealthHandler(db *pgxpool.Pool) *HealthHandler{
	return &HealthHandler{
		db: db,
	}
}

//http handler
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	postgresStatus := "ok"

	if err := h.db.Ping(r.Context()); err != nil {
		postgresStatus = "down"
	}
	
	
	response := HealthResponse{
		Status: "ok",
		Postgres: postgresStatus,
	}

	writeJSON(w, http.StatusOK, response)
}

// function helper 
func writeJSON(w http.ResponseWriter, status int, payload any) {
	// transformer une valeur Go en JSON.
	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write(data); err != nil {
		return
	}
}