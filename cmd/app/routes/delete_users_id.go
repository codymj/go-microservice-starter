package routes

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// deleteUsersId handles request to DELETE /users/{id}
func (h *handler) deleteUsersId(w http.ResponseWriter, r *http.Request) {
	// parse id from path
	idParam := mux.Vars(r)["id"]
	id, err := uuid.Parse(idParam)
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// call business service to get users by id
	err = h.UserService.Delete(r.Context(), id)
	if err != nil {
		// some other error
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// write response
	w.Header().Set(_contentType, _jsonHeader)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(nil)
}
