package users

import (
	"go-microservice-starter/internal/users/users_dao"
)

const (
	InfoBeginDeleteUserById   = "begin delete user by id"
	InfoEndDeleteUserById     = "finished delete user by id"
	InfoBeginGetUserById      = "begin getting user by id"
	InfoEndGetUserById        = "end getting user by id"
	InfoBeginGetUsersByParams = "begin getting users by params"
	InfoEndGetUsersByParams   = "finished getting users by params"
	InfoBeginSaveUser         = "begin saving user"
	InfoEndSaveUser           = "finished saving user"
	InfoBeginUpdateUserById   = "begin updating user by id"
	InfoEndUpdateUserById     = "finished updating user by id"
)

// postUsersRequestToDAO transforms the POST /users DTO into the DAO user model
func postUsersRequestToDAO(r PostUsersRequest) users_dao.User {
	return users_dao.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
