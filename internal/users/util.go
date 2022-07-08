package users

import (
	"go-microservice-starter/internal/users/users_dao"
)

// transformPostUserRequest transforms the POST /users DTO into the repository
// schema for users
func transformPostUserRequest(r PostUsersRequest) users_dao.User {
	return users_dao.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}

// transformPutUserRequest transforms the PUT /users DTO into the repository
// schema for users
func transformPutUserRequest(r PutUsersRequest) users_dao.User {
	return users_dao.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
