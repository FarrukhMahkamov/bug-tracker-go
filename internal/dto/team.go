package dto

type Team struct {
	Id       int16  `json:"id" db:"id"`
	TeamName string `json:"team_name" db:"team_name"`
}

type TeamUpdate struct {
	TeamName string `json:"team_name" db:"team_name"`
}
