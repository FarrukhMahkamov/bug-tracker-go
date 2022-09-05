package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/datastruct"
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type JobType interface {
	StoreJobType(JobType datastruct.JobType) (dto.JobType, error)
	GetAllJobType() ([]datastruct.JobType, error)
	ShowJobType(JobTypeId int) (dto.JobType, error)
	DeleteJobType(JobTypeId int) error
	UpdatedJobType(JobType dto.JobTypeUpdate, JobTypeId int) error
}

type Team interface {
	StoreTeam(Team dto.Team) (dto.Team, error)
	GetAllTeam() ([]dto.Team, error)
	ShowTeam(TeamId int) (dto.Team, error)
	DeleteTeam(TeamId int) error
	UpdatedTeam(Team dto.TeamUpdate, TeamId int) error
}

type Repository struct {
	Authorization
	JobType
	Team
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		JobType: NewJobTypePostgers(db),
		Team:    NewTeamPostgers(db),
	}
}
