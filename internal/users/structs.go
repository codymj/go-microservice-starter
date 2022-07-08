package users

// PostUsersRequest models request to POST /users
type PostUsersRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// PutUsersRequest models request to PUT /users
type PutUsersRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
