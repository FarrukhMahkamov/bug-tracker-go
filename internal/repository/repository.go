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

type Repository struct {
	Authorization
	JobType
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		JobType: NewJobTypePostgers(db),
	}
}
