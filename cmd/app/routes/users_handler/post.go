package users_handler

import (
	"encoding/json"
	"fmt"
	"go-microservice-starter/cmd/app/util"
	"go-microservice-starter/internal/users"
	"io/ioutil"
	"net/http"
)

// post handles request to POST /users
func (h *handler) post(w http.ResponseWriter, r *http.Request) {
	// parse body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("failed to read body: %v", err)
		util.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// validate payload
	errors, err := h.services.ValidatorService.ValidatePostUsers(r.Context(), body)
	if err != nil {
		util.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	if errors != nil {
		util.WriteErrorResponse(w, fmt.Errorf("%s", errors), http.StatusBadRequest)
		return
	}

	// call business service to save the users
	var req users.PostUsersRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		util.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	res, err := h.services.UserService.Save(r.Context(), req)
	if err != nil {
		util.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// write response
	b, _ := json.Marshal(res)
	w.Header().Set(util.ContentType, util.JsonHeader)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(b)
}
