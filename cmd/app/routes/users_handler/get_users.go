package users_handler

import (
	"encoding/json"
	"go-microservice-starter/cmd/app/util"
	"net/http"
	"strings"
)

// getUsers handles request to GET /users
func (h *handler) getUsers(w http.ResponseWriter, r *http.Request) {
	// parse params from path
	u, _ := r.URL.Parse(r.URL.String())
	params := make(map[string]string)
	if !strings.EqualFold("", u.RawQuery) {
		params = util.ParseQueryString(u.RawQuery)
	}

	// call business service to get users
	res, err := h.services.UserService.GetByParams(r.Context(), params)

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
