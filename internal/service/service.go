package service

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/datastruct"
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

type Authortization interface {
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

type Bug interface {
	StoreBug(Bug dto.Bug) (dto.Bug, error)
	GetAllBug() ([]dto.AllBugs, error)
	CloseIssue(BugId int) error
	AddTag(Tags dto.BugTag, BugId int) error
	GetTagsByBugId(BugId int) ([]dto.Tag, error)
}
type Service struct {
	Authortization
	JobType
	Team
	Tag
	Category
	Status
	Bug
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		JobType:  NewJobTypeService(repos.JobType),
		Team:     NewTeamService(repos.Team),
		Tag:      NewTagService(repos.Tag),
		Category: NewCategoryService(repos.Category),
		Status:   NewStatusService(repos.Status),
		Bug:      NewBugService(repos.Bug),
	}
}
