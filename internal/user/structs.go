package user

// PostUserRequest models request to POST /user
type PostUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}