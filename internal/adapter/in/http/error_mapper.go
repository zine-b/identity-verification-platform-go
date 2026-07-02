package httpin

import (
	"errors"
	"net/http"

	"github.com/zineb-b/identity-verification-platform-go/internal/application/apperror"
)

func writeAppError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, apperror.ErrEmailRequired),
		errors.Is(err, apperror.ErrInvalidEmail),
		errors.Is(err, apperror.ErrPasswordRequired),
		errors.Is(err, apperror.ErrPasswordTooWeak),
		errors.Is(err, apperror.ErrPasswordTooShort):
		writeError(w, http.StatusBadRequest, err.Error())

	case errors.Is(err, apperror.ErrUserAlreadyExists):
		writeError(w, http.StatusConflict, err.Error())

	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}