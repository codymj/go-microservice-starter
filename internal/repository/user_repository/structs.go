package user_repository

// User models a user_repository entity in the database
type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedOn int64  `json:"created_on"`
	LastLogin int64  `json:"last_login"`
}