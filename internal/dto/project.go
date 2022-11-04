package dto

type Project struct {
	Id          int    `json:"id" db:"id"`
	ProjectName string `json:"project_name" db:"project_name"`
}

type ProjectUpdate struct {
	Id          int    `json:"id" db:"id"`
	ProjectName string `json:"project_name" db:"project_name"`
}

type ProjectUser struct {
	UserId []int `json:"user_id"`
}

type ProjectTeam struct {
	TeamId []int `json:"team_id"`
}
