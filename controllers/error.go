package controllers

import (
	"net/http"

	"github.com/abibby/validate"
	"github.com/abibby/validate/handler"
)

func errorResponse(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if _, ok := err.(*validate.ValidationError); ok {
		status = http.StatusUnprocessableEntity
	}

	handler.
		NewJSONResponse(map[string]string{
			"error": err.Error(),
		}).
		SetStatus(status).
		Respond(w)
}
