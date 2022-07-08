package users

import "go-microservice-starter/internal/repository/user_repository"

// transformPostUserRequest transforms the POST /users DTO into the repository
// schema for users
func transformPostUserRequest(r PostUserRequest) user_repository.User {
	return user_repository.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}

// transformPutUserRequest transforms the PUT /users DTO into the repository
// schema for users
func transformPutUserRequest(r PutUserRequest) user_repository.User {
	return user_repository.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
