package users_handler

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go-microservice-starter/cmd/app/util"
	"net/http"
)

// deleteUsersId handles request to DELETE /users/{id}
func (h *handler) deleteUsersId(w http.ResponseWriter, r *http.Request) {
	// parse id from path
	idParam := mux.Vars(r)["id"]
	id, err := uuid.Parse(idParam)
	if err != nil {
		util.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// call business service to get users by id
	err = h.services.UserService.DeleteById(r.Context(), id)
	if err != nil {
		// some other error
		util.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// write response
	w.Header().Set(util.ContentType, util.JsonHeader)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(nil)
}
