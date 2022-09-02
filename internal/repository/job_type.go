package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/datastruct"
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type JobTypePostgres struct {
	db *sqlx.DB
}

func NewJobTypePostgers(db *sqlx.DB) *JobTypePostgres {
	return &JobTypePostgres{db: db}
}

func (r *JobTypePostgres) StoreJobType(JobType datastruct.JobType) (dto.JobType, error) {
	var StoredJobType dto.JobType

	row := r.db.QueryRow(query.StoreJobType, JobType.JobTypeName)

	if err := row.Scan(&StoredJobType.Id, &StoredJobType.JobTypeName); err != nil {
		return dto.JobType{}, err
	}

	return StoredJobType, nil
}

func (r *JobTypePostgres) GetAllJobType() ([]datastruct.JobType, error) {
	var JobTypes []datastruct.JobType

	err := r.db.Select(&JobTypes, query.GetAllJobTypes)

	return JobTypes, err
}

func (r *JobTypePostgres) ShowJobType(JobTypeId int) (dto.JobType, error) {
	var JobType dto.JobType

	err := r.db.Get(&JobType, query.ShowJobType, JobTypeId)

	return JobType, err
}
