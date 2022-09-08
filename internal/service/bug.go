package service

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

type BugService struct {
	repo repository.Bug
}

func NewBugService(repo repository.Bug) *BugService {
	return &BugService{repo: repo}
}

func (s *BugService) StoreBug(Bug dto.Bug) (dto.Bug, error) {
	return s.repo.StoreBug(Bug)
}

func (s *BugService) GetAllBug() ([]dto.Bug, error) {
	return s.repo.GetAllBug()
}

func (s *BugService) CloseIssue(BugId int) error {
	return s.repo.CloseIssue(BugId)
}

func (s *BugService) AddTag(Tags dto.BugTag, BugId int) error {
	return s.repo.AddTag(Tags, BugId)
}
