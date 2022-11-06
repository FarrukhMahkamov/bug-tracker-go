package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type ProjectPostgres struct {
	db *sqlx.DB
}

func NewProjectPostgres(db *sqlx.DB) *ProjectPostgres {
	return &ProjectPostgres{db: db}
}

func (r *ProjectPostgres) StoreProject(Project dto.Project) (dto.Project, error) {
	var StoredProject dto.Project

	row := r.db.QueryRow(query.StoreProject, Project.ProjectName)

	if err := row.Scan(&StoredProject.Id, &StoredProject.ProjectName); err != nil {
		return dto.Project{}, err
	}

	return StoredProject, nil
}

func (r *ProjectPostgres) GetAllProject() ([]dto.Project, error) {
	var Projects []dto.Project

	err := r.db.Select(&Projects, query.GetAllProjects)

	return Projects, err
}

func (r *ProjectPostgres) ShowProject(ProjectId int) (dto.Project, error) {
	var Project dto.Project

	err := r.db.Get(&Project, query.ShowProject, ProjectId)

	return Project, err
}

func (r *ProjectPostgres) DeleteProject(ProjectId int) error {
	_, err := r.db.Exec(query.DeleteProject, ProjectId)

	return err
}

func (r *ProjectPostgres) UpdatedProject(Project dto.ProjectUpdate, ProjectId int) error {
	_, err := r.db.Exec(query.UpdatedProject, Project.ProjectName, ProjectId)

	return err
}

func (r *ProjectPostgres) AddUserToProject(ProjectId int, Users dto.ProjectUser) error {
	for _, users := range Users.UserId {
		_, err := r.db.Exec(query.AddUserToProject, ProjectId, users)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ProjectPostgres) AddTeamToProject(TeamId int, ProjectId int) error {
	var TeamUsers []int

	err := r.db.Select(&TeamUsers, query.SelectUsersIdFromTeamsUsers, TeamId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query.AddTeamToProject, ProjectId, TeamId)
	if err != nil {
		return err
	}

	for _, users := range TeamUsers {
		_, err := r.db.Exec(query.AddUserToProject, ProjectId, users)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ProjectPostgres) GetBugsByProjectId(UserId int, ProjectId int) ([]dto.AllBugs, error) {
	var Bugs []dto.AllBugs

	err := r.db.Select(&Bugs, query.GetBugsByProjectId, UserId, ProjectId)

	return Bugs, err
}

func (r *ProjectPostgres) GetAttachedProjects(UserId int) ([]dto.Project, error) {
	var Projects []dto.Project

	err := r.db.Select(&Projects, query.GetProjectThatAttacedUser, UserId)

	return Projects, err
}
