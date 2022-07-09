package users_dao

import "github.com/google/uuid"

// User models a user entity in the database
type User struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password,omitempty"`
	Email      string    `json:"email"`
	IsVerified bool      `json:"isVerified"`
	CreatedOn  int64     `json:"createdOn"`
	UpdatedOn  int64     `json:"updatedOn"`
}
