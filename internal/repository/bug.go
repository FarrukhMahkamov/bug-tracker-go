package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type BugPostgres struct {
	db *sqlx.DB
}

func NewBugPostgers(db *sqlx.DB) *BugPostgres {
	return &BugPostgres{db: db}
}

func (r *BugPostgres) StoreBug(Bug dto.Bug) (dto.Bug, error) {
	var StoredBug dto.Bug

	row := r.db.QueryRow(query.StoreBug, Bug.BugTitle, Bug.BugDescription, Bug.StatusId, Bug.CategoryId)

	if err := row.Scan(
		&StoredBug.Id,
		&StoredBug.BugTitle,
		&StoredBug.BugDescription,
		&StoredBug.IsCompleted,
		&StoredBug.StatusId,
		&StoredBug.CategoryId,
	); err != nil {
		return dto.Bug{}, err
	}

	return StoredBug, nil
}

func (r *BugPostgres) GetAllBug() ([]dto.Bug, error) {
	var Bugs []dto.Bug

	err := r.db.Select(&Bugs, query.GetAllBugs)

	return Bugs, err
}

func (r *BugPostgres) CloseIssue(BugId int) error {
	_, err := r.db.Exec(query.CloseIssue, BugId)

	return err
}

func (r *BugPostgres) AddTag(Tags dto.BugTag, BugId int) error {
	for _, tags := range Tags.TagId {
		_, err := r.db.Exec(query.AddTag, BugId, tags)
		if err != nil {
			return err
		}
	}
	return nil
}
