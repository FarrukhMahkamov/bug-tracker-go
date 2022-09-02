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
