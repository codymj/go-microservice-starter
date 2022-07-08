package routes

import (
	"encoding/json"
	"net/http"
	"strings"
)

// getUsers handles request to GET /users
func (h *handler) getUsers(w http.ResponseWriter, r *http.Request) {
	// parse params from path
	u, _ := r.URL.Parse(r.URL.String())
	params := make(map[string]string)
	if !strings.EqualFold("", u.RawQuery) {
		params = parseQueryString(u.RawQuery)
	}

	// call business service to get users
	var res any
	var err error

	if len(params) == 0 {
		// no params set, do a get all
		res, err = h.UserService.GetAll(r.Context())
	} else {
		// at least one param set
		res, err = h.UserService.GetByParams(r.Context(), params)
	}

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
