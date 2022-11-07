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

func (s *BugService) GetAllBug() ([]dto.AllBugs, error) {
	return s.repo.GetAllBug()
}

func (s *BugService) CloseIssue(BugId int) error {
	return s.repo.CloseIssue(BugId)
}

func (s *BugService) AddTag(Tags dto.BugTag, BugId int) error {
	return s.repo.AddTag(Tags, BugId)
}

func (s *BugService) GetTagsByBugId(BugId int) ([]dto.Tag, error) {
	return s.repo.GetTagsByBugId(BugId)
}

func (s *BugService) AttachUserToBug(BugId int, Users dto.AttachUsers) error {
	return s.repo.AttachUserToBug(BugId, Users)
}

func (s *BugService) DeattachUserFromBug(BugId int, Users dto.AttachUsers) error {
	return s.repo.DeattachUserFromBug(BugId, Users)
}

func (s *BugService) AttachTeamToBug(BugId int, TeamId int) error {
	return s.repo.AttachTeamToBug(BugId, TeamId)
}

func (s *BugService) DetachTeamFromBug(BugId int, TeamId int) error {
	return s.repo.DetachTeamFromBug(BugId, TeamId)
}
