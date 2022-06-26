package routes

import (
	"encoding/json"
	"fmt"
	"go-microservice-starter/internal/greeting"
	"io/ioutil"
	"net/http"
)

// postGreeting handles request to POST /greeting
func (h *handler) postGreeting(w http.ResponseWriter, r *http.Request) {
	// parse body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("failed to read body: %v", err)
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// validate payload
	errors, err := h.ValidatorService.ValidatePostHello(r.Context(), payload)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	if errors != nil {
		writeErrorResponse(w, fmt.Errorf("%s", errors), http.StatusBadRequest)
		return
	}

	// pass to service
	var request greeting.PostRequest
	err = json.Unmarshal(payload, &request)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	res := h.GreetingService.SayHello(r.Context(), request)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
	}

	// write response
	b, _ := json.Marshal(res)
	w.Header().Set(_contentType, _jsonHeader)
	_, _ = w.Write(b)
}
