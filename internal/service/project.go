package service

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

type ProjectService struct {
	repo repository.Project
}

func NewPorjectService(repo repository.Project) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) StoreProject(Project dto.Project) (dto.Project, error) {
	return s.repo.StoreProject(Project)
}

func (s *ProjectService) GetAllProject() ([]dto.Project, error) {
	return s.repo.GetAllProject()
}

func (s *ProjectService) ShowProject(ProjectId int) (dto.Project, error) {
	return s.repo.ShowProject(ProjectId)
}

func (s *ProjectService) DeleteProject(ProjectId int) error {
	return s.repo.DeleteProject(ProjectId)
}

func (s *ProjectService) UpdatedProject(Project dto.ProjectUpdate, ProjectId int) error {
	return s.repo.UpdatedProject(Project, ProjectId)
}

func (s *ProjectService) AddUserToProject(Users dto.ProjectUser, ProjectId int) error {
	return s.repo.AddUserToProject(ProjectId, Users)
}

func (s *ProjectService) AddTeamToProject(TeamId int, ProjectId int) error {
	return s.repo.AddTeamToProject(TeamId, ProjectId)
}

func (s *ProjectService) GetBugsByProjectId(UserId int, ProjectId int) ([]dto.AllBugs, error) {
	return s.repo.GetBugsByProjectId(UserId, ProjectId)
}

func (s *ProjectService) GetAttachedProjects(UserId int) ([]dto.Project, error) {
	return s.repo.GetAttachedProjects(UserId)
}

func (s *ProjectService) RemoveUsersFromProject(Users dto.ProjectUser, ProjectId int) error {
	return s.repo.RemoveUsersFromProject(Users, ProjectId)
}
