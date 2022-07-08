package routes

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

// getUsersId handles request to GET /users/{id}
func (h *handler) getUsersId(w http.ResponseWriter, r *http.Request) {
	// parse id from path
	strId := path.Base(r.URL.Path)
	id, err := strconv.Atoi(strId)
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// call business service to get users by id
	res, err := h.UserService.GetById(r.Context(), int64(id))
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
