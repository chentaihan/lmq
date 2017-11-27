package server

import (
	"net/http"
	"strings"
)

// apiError is an error wrapper that also
// holds information about response status codes.
type apiError struct {
	error
	statusCode int
}

// HTTPErrorStatusCode returns a status code.
func (e apiError) HTTPErrorStatusCode() int {
	return e.statusCode
}

// NewErrorWithStatusCode allows you to associate
// a specific HTTP Status Code to an error.
// The server will take that code and set
// it as the response status.
func NewErrorWithStatusCode(err error, code int) error {
	return apiError{err, code}
}

// NewBadRequestError creates a new API error
// that has the 400 HTTP status code associated to it.
func NewBadRequestError(err error) error {
	return NewErrorWithStatusCode(err, http.StatusBadRequest)
}

// NewRequestForbiddenError creates a new API error
// that has the 403 HTTP status code associated to it.
func NewRequestForbiddenError(err error) error {
	return NewErrorWithStatusCode(err, http.StatusForbidden)
}

// NewRequestNotFoundError creates a new API error
// that has the 404 HTTP status code associated to it.
func NewRequestNotFoundError(err error) error {
	return NewErrorWithStatusCode(err, http.StatusNotFound)
}

// NewRequestConflictError creates a new API error
// that has the 409 HTTP status code associated to it.
func NewRequestConflictError(err error) error {
	return NewErrorWithStatusCode(err, http.StatusConflict)
}

func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
}

func MakeErrorHandler(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode := GetHTTPErrorStatusCode(err)
		http.Error(w, err.Error(), statusCode)
	}
}

func GetHTTPErrorStatusCode(err error) int {
	if err == nil {
		return http.StatusInternalServerError
	}
	var statusCode int
	errMsg := err.Error()
	errStr := strings.ToLower(errMsg)
	for _, status := range []struct {
		keyword string
		code    int
	}{
		{"not found", http.StatusNotFound},
		{"cannot find", http.StatusNotFound},
		{"no such", http.StatusNotFound},
		{"bad parameter", http.StatusBadRequest},
		{"no command", http.StatusBadRequest},
		{"conflict", http.StatusConflict},
		{"impossible", http.StatusNotAcceptable},
		{"wrong login/password", http.StatusUnauthorized},
		{"unauthorized", http.StatusUnauthorized},
		{"hasn't been activated", http.StatusForbidden},
		{"this node", http.StatusServiceUnavailable},
		{"needs to be unlocked", http.StatusServiceUnavailable},
		{"certificates have expired", http.StatusServiceUnavailable},
		{"repository does not exist", http.StatusNotFound},
	} {
		if strings.Contains(errStr, status.keyword) {
			statusCode = status.code
			break
		}
	}
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	return statusCode
}