package dto

type Category struct {
	Id           int16  `json:"id" db:"id"`
	CategoryName string `json:"category_name" db:"category_name"`
}

type CategoryUpdate struct {
	CategoryName string `json:"category_name" db:"category_name"`
}
