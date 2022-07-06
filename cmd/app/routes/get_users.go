package routes

import (
	"encoding/json"
	"net/http"
)

// getUsers handles request to GET /users
func (h *handler) getUsers(w http.ResponseWriter, r *http.Request) {
	// parse body
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	err = fmt.Errorf("failed to read body: %v", err)
	//	writeErrorResponse(w, err, http.StatusInternalServerError)
	//	return
	//}

	// validate payload
	//errors, err := h.ValidatorService.ValidateGetUsers(r.Context(), body)
	//if err != nil {
	//	writeErrorResponse(w, err, http.StatusInternalServerError)
	//	return
	//}
	//if errors != nil {
	//	writeErrorResponse(w, fmt.Errorf("%s", errors), http.StatusBadRequest)
	//	return
	//}

	// pass to service
	//var request user.PostGreetingRequest
	//err = json.Unmarshal(body, &request)
	//if err != nil {
	//	writeErrorResponse(w, err, http.StatusInternalServerError)
	//	return
	//}
	res, err := h.UserService.GetAll(r.Context())
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
	}

	// write response
	b, _ := json.Marshal(res)
	w.Header().Set(_contentType, _jsonHeader)
	_, _ = w.Write(b)
}
