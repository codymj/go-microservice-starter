package users

import (
	"go-microservice-starter/internal/users/users_dao"
)

// postUsersRequestToDAO transforms the POST /users DTO into the DAO user model
func postUsersRequestToDAO(r PostUsersRequest) users_dao.User {
	return users_dao.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
