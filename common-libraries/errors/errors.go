package errors

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Msg  string
	Code int
}

func (e Error) Error() string {
	return e.Msg
}

func NewNotFound(message string) *Error {
	return &Error{message, 404}
}

func NewBadRequest(message string) *Error {
	return &Error{message, 400}
}

func NewUnauthorized(message string) *Error {
	return &Error{message, 401}
}

func NewForbidden(message string) *Error {
	return &Error{message, 403}
}

func NewInternal(message string) *Error {
	return &Error{message, 500}
}

func HandleError(w http.ResponseWriter, err *Error) {
	errJSON := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)

	data, marshalErr := json.Marshal(errJSON)
	if marshalErr != nil {
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	_, writeErr := w.Write(data)
	if writeErr != nil {
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	return
}
