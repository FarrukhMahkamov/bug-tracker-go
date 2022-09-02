package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/datastruct"
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Authortization interface {
}

type JobType interface {
	StoreJobType(JobType datastruct.JobType) (dto.JobType, error)
	GetAllJobType() ([]datastruct.JobType, error)
	ShowJobType(JobTypeId int) (dto.JobType, error)
}

type Repository struct {
	Authortization
	JobType
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		JobType: NewJobTypePostgers(db),
	}
}
