package dto

type User struct {
	Id       int64   `json:"id" db:"id"`
	Name     string  `json:"name" db:"name" binding:"required"`
	UserName string  `json:"username" db:"username" binding:"required"`
	Email    string  `json:"email" db:"email" binding:"required"`
	Password string  `json:"password" db:"password" binding:"required"`
	Image    *string `json:"image" db:"image"`
}

type UserForUi struct {
	Id       int64   `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	UserName string  `json:"username" db:"username"`
	Email    string  `json:"email" db:"email"`
	Image    *string `json:"image" db:"image"`
	Token    *string `json:"token"`
}


type UserSignInFields struct {
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
