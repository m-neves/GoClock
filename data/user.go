package data

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserId   int    `json:"user_id,omitempty"`
}
