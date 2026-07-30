package errors

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(
	w http.ResponseWriter,
	status int,
	data any,
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func WriteError(
	w http.ResponseWriter,
	err error,
) {

	// Our custom error
	if apiErr, ok := err.(APIError); ok {

		WriteJSON(
			w,
			apiErr.Status,
			apiErr,
		)

		return
	}

	// Unknown error
	WriteJSON(
		w,
		http.StatusInternalServerError,
		APIError{
			Code:    CodeInternal,
			Message: err.Error(),
		},
	)
}
