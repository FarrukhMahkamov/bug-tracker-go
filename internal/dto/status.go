package dto

type Status struct {
	Id         int16  `json:"id" db:"id"`
	StatusName string `json:"status_name" db:"status_name"`
}

type StatusUpdate struct {
	StatusName string `json:"status_name" db:"status_name"`
}
