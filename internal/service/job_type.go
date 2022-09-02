package service

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/datastruct"
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

type JobTypeService struct {
	repo repository.JobType
}

func NewJobTypeService(repo repository.JobType) *JobTypeService {
	return &JobTypeService{repo: repo}
}

func (s *JobTypeService) StoreJobType(JobType datastruct.JobType) (dto.JobType, error) {
	return s.repo.StoreJobType(JobType)
}

func (s *JobTypeService) GetAllJobType() ([]datastruct.JobType, error) {
	return s.repo.GetAllJobType()
}
