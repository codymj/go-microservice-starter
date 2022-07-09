package routes

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go-microservice-starter/internal/users"
	"io/ioutil"
	"net/http"
)

// putUsersId handles request to PUT /users/{id}
func (h *handler) putUsersId(w http.ResponseWriter, r *http.Request) {
	// parse id from path
	idParam := mux.Vars(r)["id"]
	id, err := uuid.Parse(idParam)
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// parse body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("failed to read body: %v", err)
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// validate payload
	errors, err := h.ValidatorService.ValidatePutUsersId(r.Context(), body)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	if errors != nil {
		writeErrorResponse(w, fmt.Errorf("%s", errors), http.StatusBadRequest)
		return
	}

	// call business service to update the users
	var req users.PutUsersIdRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	res, err := h.UserService.UpdateById(r.Context(), id, req)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// write response
	b, _ := json.Marshal(res)
	w.Header().Set(_contentType, _jsonHeader)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}
