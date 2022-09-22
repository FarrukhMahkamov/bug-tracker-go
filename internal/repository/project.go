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
	var Categories []dto.Project

	err := r.db.Select(&Categories, query.GetAllCategories)

	return Categories, err
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
