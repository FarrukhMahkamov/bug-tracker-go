package service

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

type TeamService struct {
	repo repository.Team
}

func NewTeamService(repo repository.Team) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) StoreTeam(Team dto.Team) (dto.Team, error) {
	return s.repo.StoreTeam(Team)
}

func (s *TeamService) GetAllTeam() ([]dto.Team, error) {
	return s.repo.GetAllTeam()
}

func (s *TeamService) ShowTeam(TeamId int) (dto.Team, error) {
	return s.repo.ShowTeam(TeamId)
}

func (s *TeamService) DeleteTeam(TeamId int) error {
	return s.repo.DeleteTeam(TeamId)
}

func (s *TeamService) UpdatedTeam(Team dto.TeamUpdate, TeamId int) error {
	return s.repo.UpdatedTeam(Team, TeamId)
}

func (s *TeamService) AddUsersToTeam(TeamId int, Users dto.TeamUsers) error {
	return s.repo.AddUsersToTeam(TeamId, Users)
}

func (s *TeamService) GetTeamUsers(TeamId int) ([]dto.UserForUi, error) {
	return s.repo.GetTeamUsers(TeamId)
}

func (s *TeamService) RemoveUsersFromTeam(TeamId int, Users dto.TeamUsers) error {
	return s.repo.RemoveUsersFromTeam(TeamId, Users)
}
