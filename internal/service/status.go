package service

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

type StatusService struct {
	repo repository.Status
}

func NewStatusService(repo repository.Status) *StatusService {
	return &StatusService{repo: repo}
}

func (s *StatusService) StoreStatus(Status dto.Status) (dto.Status, error) {
	return s.repo.StoreStatus(Status)
}

func (s *StatusService) GetAllStatus() ([]dto.Status, error) {
	return s.repo.GetAllStatus()
}

func (s *StatusService) ShowStatus(StatusId int) (dto.Status, error) {
	return s.repo.ShowStatus(StatusId)
}

func (s *StatusService) DeleteStatus(StatusId int) error {
	return s.repo.DeleteStatus(StatusId)
}

func (s *StatusService) UpdatedStatus(Status dto.StatusUpdate, StatusId int) error {
	return s.repo.UpdatedStatus(Status, StatusId)
}
