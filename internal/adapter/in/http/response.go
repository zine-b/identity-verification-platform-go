package httpin

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	data = append(data, '\n')

	w.WriteHeader(status)

	if _, err := w.Write(data); err != nil {
		return
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}
