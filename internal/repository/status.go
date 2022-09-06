package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type StatusPostgres struct {
	db *sqlx.DB
}

func NewStatusPostgers(db *sqlx.DB) *StatusPostgres {
	return &StatusPostgres{db: db}
}

func (r *StatusPostgres) StoreStatus(Status dto.Status) (dto.Status, error) {
	var StoredStatus dto.Status

	row := r.db.QueryRow(query.StoreStatus, Status.StatusName)

	if err := row.Scan(&StoredStatus.Id, &StoredStatus.StatusName); err != nil {
		return dto.Status{}, err
	}

	return StoredStatus, nil
}

func (r *StatusPostgres) GetAllStatus() ([]dto.Status, error) {
	var Statuses []dto.Status

	err := r.db.Select(&Statuses, query.GetAllStatuses)

	return Statuses, err
}

func (r *StatusPostgres) ShowStatus(StatusId int) (dto.Status, error) {
	var Status dto.Status

	err := r.db.Get(&Status, query.ShowStatus, StatusId)

	return Status, err
}

func (r *StatusPostgres) DeleteStatus(StatusId int) error {
	_, err := r.db.Exec(query.DeleteStatus, StatusId)

	return err
}

func (r *StatusPostgres) UpdatedStatus(Status dto.StatusUpdate, StatusId int) error {
	_, err := r.db.Exec(query.UpdatedStatus, Status.StatusName, StatusId)

	return err
}
