package users

import (
	"go-microservice-starter/internal/users/users_dao"
)

// postUsersReqToDAO transforms the POST /users DTO into the DAO user model
func postUsersReqToDAO(r PostUsersRequest) users_dao.User {
	return users_dao.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
