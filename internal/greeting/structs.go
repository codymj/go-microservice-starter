package greeting

// PostRequest models request to POST /greeting
type PostRequest struct {
	Name string `json:"name"`
}

// PostResponse models response from POST /greeting
type PostResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
