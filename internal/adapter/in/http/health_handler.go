package http

import (
	"net/http"
	"encoding/json"
)
type HealthHandler struct{}

func NewHealthHandler() *HealthHandler{
	return &HealthHandler{}
}

//http handler (fct)
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}