package routes

import (
	"encoding/json"
	"fmt"
	"go-microservice-starter/internal/users"
	"io/ioutil"
	"net/http"
)

// postUsers handles request to POST /users
func (h *handler) postUsers(w http.ResponseWriter, r *http.Request) {
	// parse body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("failed to read body: %v", err)
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// validate payload
	errors, err := h.ValidatorService.ValidatePostUsers(r.Context(), body)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	if errors != nil {
		writeErrorResponse(w, fmt.Errorf("%s", errors), http.StatusBadRequest)
		return
	}

	// call business service to save the users
	var req users.PostUsersRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	res, err := h.UserService.Save(r.Context(), req)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// write response
	b, _ := json.Marshal(res)
	w.Header().Set(_contentType, _jsonHeader)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(b)
}
