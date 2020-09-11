package http

import (
	"fmt"
	"net/http"
)

// RestError defines the error response of a REST request
type RestError struct {
	Status           int    `json:"status"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

// Asserts that RestError implements error
var _ error = (*RestError)(nil)

func (r *RestError) Error() string {
	return fmt.Sprintf("HTTP %d - %s", r.Status, r.Message)
}

func BadRequest(message string) *RestError {
	return &RestError{Status: http.StatusBadRequest, Message: message}
}

func Forbidden(message string) *RestError {
	return &RestError{Status: http.StatusForbidden, Message: message}
}

func Unauthorized(message string) *RestError {
	return &RestError{Status: http.StatusUnauthorized, Message: message}
}

func NotFound(message string) *RestError {
	return &RestError{Status: http.StatusNotFound, Message: message}
}

func InternalServerError(err error) *RestError {
	return &RestError{
		Status:           http.StatusInternalServerError,
		Message:          "Internal Server Error",
		DeveloperMessage: err.Error(),
	}
}
