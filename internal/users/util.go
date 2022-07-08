package users

import "go-microservice-starter/internal/repository/users_repository"

// transformPostUserRequest transforms the POST /users DTO into the repository
// schema for users
func transformPostUserRequest(r PostUsersRequest) users_repository.User {
	return users_repository.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}

// transformPutUserRequest transforms the PUT /users DTO into the repository
// schema for users
func transformPutUserRequest(r PutUsersRequest) users_repository.User {
	return users_repository.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
