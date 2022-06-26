package greeting

// PostGreetingRequest models request to POST /greeting
type PostGreetingRequest struct {
	Name string `json:"name"`
}

// PostGreetingResponse models response from POST /greeting
type PostGreetingResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
