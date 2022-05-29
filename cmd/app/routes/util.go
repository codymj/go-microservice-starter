package routes

import (
	"encoding/json"
	"net/http"
)

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
