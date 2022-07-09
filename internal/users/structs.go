package users

// PostUsersRequest models request to POST /users
type PostUsersRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// PutUsersIdRequest models request to PUT /users/{id}
type PutUsersIdRequest struct {
	Email string `json:"email"`
}
