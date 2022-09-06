package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type TagPostgres struct {
	db *sqlx.DB
}

func NewTagPostgers(db *sqlx.DB) *TagPostgres {
	return &TagPostgres{db: db}
}

func (r *TagPostgres) StoreTag(Tag dto.Tag) (dto.Tag, error) {
	var StoredTag dto.Tag

	row := r.db.QueryRow(query.StoreTag, Tag.TagName)

	if err := row.Scan(&StoredTag.Id, &StoredTag.TagName); err != nil {
		return dto.Tag{}, err
	}

	return StoredTag, nil
}

func (r *TagPostgres) GetAllTag() ([]dto.Tag, error) {
	var Tags []dto.Tag

	err := r.db.Select(&Tags, query.GetAllTags)

	return Tags, err
}

func (r *TagPostgres) ShowTag(TagId int) (dto.Tag, error) {
	var Tag dto.Tag

	err := r.db.Get(&Tag, query.ShowTag, TagId)

	return Tag, err
}

func (r *TagPostgres) DeleteTag(TagId int) error {
	_, err := r.db.Exec(query.DeleteTag, TagId)

	return err
}

func (r *TagPostgres) UpdatedTag(Tag dto.TagUpdate, TagId int) error {
	_, err := r.db.Exec(query.UpdatedTag, Tag.TagName, TagId)

	return err
}
