package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type TeamPostgres struct {
	db *sqlx.DB
}

func NewTeamPostgers(db *sqlx.DB) *TeamPostgres {
	return &TeamPostgres{db: db}
}

func (r *TeamPostgres) StoreTeam(Team dto.Team) (dto.Team, error) {
	var StoredTeam dto.Team

	row := r.db.QueryRow(query.StoreTeam, Team.TeamName)

	if err := row.Scan(&StoredTeam.Id, &StoredTeam.TeamName); err != nil {
		return dto.Team{}, err
	}

	return StoredTeam, nil
}

func (r *TeamPostgres) GetAllTeam() ([]dto.Team, error) {
	var Teams []dto.Team

	err := r.db.Select(&Teams, query.GetAllTeams)

	return Teams, err
}

func (r *TeamPostgres) ShowTeam(TeamId int) (dto.Team, error) {
	var Team dto.Team

	err := r.db.Get(&Team, query.ShowTeam, TeamId)

	return Team, err
}

func (r *TeamPostgres) DeleteTeam(TeamId int) error {
	_, err := r.db.Exec(query.DeleteTeam, TeamId)

	return err
}

func (r *TeamPostgres) UpdatedTeam(Team dto.TeamUpdate, TeamId int) error {
	_, err := r.db.Exec(query.UpdatedTeam, Team.TeamName, TeamId)

	return err
}