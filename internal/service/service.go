package service

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/datastruct"
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

type Authortization interface {
}

type JobType interface {
	StoreJobType(JobType datastruct.JobType) (dto.JobType, error)
	GetAllJobType() ([]datastruct.JobType, error)
	ShowJobType(JobTypeId int) (dto.JobType, error)
	DeleteJobType(JobTypeId int) error
	UpdatedJobType(JobType dto.JobTypeUpdate, JobTypeId int) error
}

type Service struct {
	Authortization
	JobType
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		JobType: NewJobTypeService(repos.JobType),
	}
}
