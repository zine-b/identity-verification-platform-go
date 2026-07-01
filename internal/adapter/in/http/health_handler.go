package httpin

import (
	"net/http"
	"encoding/json"
)
type HealthHandler struct{}

type HealthResponse struct{
	Status string `json:"status"`
}

func NewHealthHandler() *HealthHandler{
	return &HealthHandler{}
}

//http handler
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
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