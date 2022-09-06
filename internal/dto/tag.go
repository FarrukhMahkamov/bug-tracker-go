package dto

type Tag struct {
	Id      int16  `json:"id" db:"id"`
	TagName string `json:"tag_name" db:"tag_name"`
}

type TagUpdate struct {
	TagName string `json:"tag_name" db:"tag_name"`
}
