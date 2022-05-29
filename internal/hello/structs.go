package hello

// PostRequest models request to POST /hello
type PostRequest struct {
	Name string `json:"name"`
}

// PostResponse models response from POST /hello
type PostResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
