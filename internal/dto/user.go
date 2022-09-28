package dto

type User struct {
	Id       int64   `json:"id" db:"id"`
	Name     string  `json:"name" db:"name" binding:"required"`
	UserName string  `json:"username" db:"username" binding:"required"`
	Email    string  `json:"email" db:"email" binding:"required"`
	Password string  `json:"password" db:"password"`
	Image    *string `json:"image" db:"image"`
}

type UserForUi struct {
	Id       int64   `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	UserName string  `json:"user_name" db:"user_naname"`
	Email    string  `json:"email" db:"email"`
	Image    *string `json:"image" db:"image"`
}
