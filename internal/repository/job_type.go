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
