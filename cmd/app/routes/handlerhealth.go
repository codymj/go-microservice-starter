package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// getHealth performs a health check on the service
func getHealth(w http.ResponseWriter, _ *http.Request) {
	response := GenericResponse{
		Status:  "ok",
		Message: "service is healthy",
	}
	b, err := json.Marshal(response)
	if err != nil {
		err = fmt.Errorf("failed to marshal health response: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(_contentType, _jsonHeader)
	_, err = w.Write(b)
}
