package routes

import (
	"encoding/json"
	"net/http"
)

// getUsers handles request to GET /users
func (h *handler) getUsers(w http.ResponseWriter, r *http.Request) {
	// call business service to get users
	res, err := h.UserService.GetAll(r.Context())
	if res == nil {
		// no users found
		w.WriteHeader(http.StatusNoContent)
		_ = json.NewEncoder(w).Encode(nil)
		return
	}
	if err != nil {
		// some other error
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// write response
	b, _ := json.Marshal(res)
	w.Header().Set(_contentType, _jsonHeader)
	_, _ = w.Write(b)
}
