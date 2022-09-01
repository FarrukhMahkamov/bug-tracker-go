package dto

type User struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	UserName string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Image    *string `json:"image"`
}
