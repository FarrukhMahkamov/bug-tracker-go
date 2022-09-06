package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgers(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) StoreCategory(Category dto.Category) (dto.Category, error) {
	var StoredCategory dto.Category

	row := r.db.QueryRow(query.StoreCategory, Category.CategoryName)

	if err := row.Scan(&StoredCategory.Id, &StoredCategory.CategoryName); err != nil {
		return dto.Category{}, err
	}

	return StoredCategory, nil
}

func (r *CategoryPostgres) GetAllCategory() ([]dto.Category, error) {
	var Categories []dto.Category

	err := r.db.Select(&Categories, query.GetAllCategories)

	return Categories, err
}

func (r *CategoryPostgres) ShowCategory(CategoryId int) (dto.Category, error) {
	var Category dto.Category

	err := r.db.Get(&Category, query.ShowCategory, CategoryId)

	return Category, err
}

func (r *CategoryPostgres) DeleteCategory(CategoryId int) error {
	_, err := r.db.Exec(query.DeleteCategory, CategoryId)

	return err
}

func (r *CategoryPostgres) UpdatedCategory(Category dto.CategoryUpdate, CategoryId int) error {
	_, err := r.db.Exec(query.UpdatedCategory, Category.CategoryName, CategoryId)

	return err
}
