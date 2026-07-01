package httpin

import (
	"encoding/json"
	"net/http"

	portin "github.com/zineb-b/identity-verification-platform-go/internal/application/port/in"
)


type AuthHandler struct {
	authUseCase portin.AuthUseCase
}

func NewAuthHandler(authUseCase portin.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var request SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	result, err := h.authUseCase.Signup(r.Context(), portin.SignupCommand{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, result)
}