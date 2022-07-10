package util

import (
	"encoding/json"
	"net/http"
	"strings"
)

// ParseQueryString manually parses query string
func ParseQueryString(s string) map[string]string {
	params := make(map[string]string)

	splitQueries := strings.Split(s, "&")
	for _, query := range splitQueries {
		keyValue := strings.Split(query, "=")
		key := keyValue[0]
		value := keyValue[1]
		params[key] = value
	}

	return params
}

// WriteErrorResponse builds a response in case of error
func WriteErrorResponse(w http.ResponseWriter, err error, code int) {
	response := GenericResponse{
		Status:  "failed",
		Message: err.Error(),
	}

	w.Header().Set(ContentType, JsonHeader)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(response)
}
