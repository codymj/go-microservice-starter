package users

import "github.com/google/uuid"

// PostUsersRequest models request to POST /users
type PostUsersRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// PutUsersIdRequest models request to PUT /users/{id}
type PutUsersIdRequest struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	IsVerified bool      `json:"isVerified"`
	CreatedOn  int64     `json:"createdOn"`
	UpdatedOn  int64     `json:"updatedOn"`
}
