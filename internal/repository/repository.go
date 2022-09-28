package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/datastruct"
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	RegisterUser(user dto.User) (dto.UserForUi, error)
}

type JobType interface {
	StoreJobType(JobType datastruct.JobType) (dto.JobType, error)
	GetAllJobType() ([]datastruct.JobType, error)
	ShowJobType(JobTypeId int) (dto.JobType, error)
	DeleteJobType(JobTypeId int) error
	UpdatedJobType(JobType dto.JobTypeUpdate, JobTypeId int) error
}

type Team interface {
	StoreTeam(Team dto.Team) (dto.Team, error)
	GetAllTeam() ([]dto.Team, error)
	ShowTeam(TeamId int) (dto.Team, error)
	DeleteTeam(TeamId int) error
	UpdatedTeam(Team dto.TeamUpdate, TeamId int) error
}

type Tag interface {
	StoreTag(Tag dto.Tag) (dto.Tag, error)
	GetAllTag() ([]dto.Tag, error)
	ShowTag(TagId int) (dto.Tag, error)
	DeleteTag(TagId int) error
	UpdatedTag(Tag dto.TagUpdate, TagId int) error
}

type Category interface {
	StoreCategory(Category dto.Category) (dto.Category, error)
	GetAllCategory() ([]dto.Category, error)
	ShowCategory(CategoryId int) (dto.Category, error)
	DeleteCategory(CategoryId int) error
	UpdatedCategory(Category dto.CategoryUpdate, CategoryId int) error
}

type Status interface {
	StoreStatus(Status dto.Status) (dto.Status, error)
	GetAllStatus() ([]dto.Status, error)
	ShowStatus(StatusId int) (dto.Status, error)
	DeleteStatus(StatusId int) error
	UpdatedStatus(Status dto.StatusUpdate, StatusId int) error
}

type Project interface {
	StoreProject(Project dto.Project) (dto.Project, error)
	GetAllProject() ([]dto.Project, error)
	ShowProject(ProjectId int) (dto.Project, error)
	DeleteProject(ProjectId int) error
	UpdatedProject(Project dto.ProjectUpdate, ProjectId int) error
}

type Bug interface {
	StoreBug(Bug dto.Bug) (dto.Bug, error)
	GetAllBug() ([]dto.AllBugs, error)
	CloseIssue(BugId int) error
	AddTag(Tags dto.BugTag, BugId int) error
	GetTagsByBugId(BugId int) ([]dto.Tag, error)
}
type Repository struct {
	Authorization
	JobType
	Team
	Tag
	Category
	Status
	Bug
	Project
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		JobType:       NewJobTypePostgers(db),
		Team:          NewTeamPostgers(db),
		Tag:           NewTagPostgers(db),
		Category:      NewCategoryPostgers(db),
		Status:        NewStatusPostgers(db),
		Bug:           NewBugPostgers(db),
		Project:       NewProjectPostgres(db),
	}
}
