package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	query "github.com/FarrukhMahkamov/bugtracker/internal/repository/query"
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

	row := r.db.QueryRow(query.StoreBug, Bug.BugTitle, Bug.BugDescription, Bug.StatusId, Bug.CategoryId, Bug.ProjectId)

	if err := row.Scan(
		&StoredBug.Id,
		&StoredBug.BugTitle,
		&StoredBug.BugDescription,
		&StoredBug.IsCompleted,
		&StoredBug.StatusId,
		&StoredBug.CategoryId,
		&StoredBug.ProjectId,
	); err != nil {
		return dto.Bug{}, err
	}

	return StoredBug, nil
}

func (r *BugPostgres) GetAllBug() ([]dto.AllBugs, error) {
	var Bugs []dto.AllBugs

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

func (r *BugPostgres) GetTagsByBugId(BugId int) ([]dto.Tag, error) {
	var Tags []dto.Tag

	err := r.db.Select(&Tags, query.GetTagsByBugId, BugId)

	return Tags, err
}

func (r *BugPostgres) AttachUserToBug(BugId int, Users dto.AttachUsers) error {
	for _, users := range Users.UserId {
		_, err := r.db.Exec(query.AttachUserToBug, BugId, users)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *BugPostgres) AttachTeamToBug(BugId int, TeamId int) error {
	var TeamUsers []int

	err := r.db.Select(&TeamUsers, query.SelectUsersIdFromTeamsUsers, TeamId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query.AttachTeamToBug, BugId, TeamId)
	if err != nil {
		return err
	}

	for _, users := range TeamUsers {
		_, err := r.db.Exec(query.AttachUserToBug, BugId, users)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *BugPostgres) DeattachUserFromBug(BugId int, Users dto.AttachUsers) error {
	for _, users := range Users.UserId {
		_, err := r.db.Exec(query.DeattachUserFromBug, BugId, users)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *BugPostgres) DetachTeamFromBug(BugId int, TeamId int) error {
	var TeamUsers []int

	err := r.db.Select(&TeamUsers, query.SelectUsersIdFromTeamsUsers, TeamId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query.DetachTeamFromBug, BugId, TeamId)
	if err != nil {
		return err
	}

	for _, users := range TeamUsers {
		_, err := r.db.Exec(query.DeattachUserFromBug, BugId, users)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *BugPostgres) UpdateBug(BugId int, UpdateBug dto.BugUpdate) error {
	_, err := r.db.Exec(query.UpdateBug, UpdateBug.BugTitle, UpdateBug.BugDescription, UpdateBug.StatusId, UpdateBug.CategoryId, BugId)

	return err
}
