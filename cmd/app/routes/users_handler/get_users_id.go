package users_handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go-microservice-starter/cmd/app/util"
	"net/http"
)

// getUsersId handles request to GET /users/{id}
func (h *handler) getUsersId(w http.ResponseWriter, r *http.Request) {
	// parse id from path
	idParam := mux.Vars(r)["id"]
	id, err := uuid.Parse(idParam)
	if err != nil {
		util.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// call business service to get users by id
	res, err := h.services.UserService.GetById(r.Context(), id)
	if res == nil {
		// no users found
		w.WriteHeader(http.StatusNoContent)
		_ = json.NewEncoder(w).Encode(nil)
		return
	}
	if err != nil {
		// some other error
		util.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// write response
	b, _ := json.Marshal(res)
	w.Header().Set(util.ContentType, util.JsonHeader)
	_, _ = w.Write(b)
}
