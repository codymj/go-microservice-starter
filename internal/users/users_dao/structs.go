package users_dao

// User models a users entity in the database
type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email"`
	CreatedOn int64  `json:"created_on"`
	LastLogin int64  `json:"last_login"`
}
