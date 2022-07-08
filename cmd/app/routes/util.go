package routes

import (
	"encoding/json"
	"net/http"
	"strings"
)

// parseQueryString manually parses query string
func parseQueryString(s string) map[string]string {
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

// writeErrorResponse builds a response in case of error
func writeErrorResponse(w http.ResponseWriter, err error, code int) {
	response := GenericResponse{
		Status:  "failed",
		Message: err.Error(),
	}

	w.Header().Set(_contentType, _jsonHeader)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(response)
}
