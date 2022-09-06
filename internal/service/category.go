package service

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) StoreCategory(Category dto.Category) (dto.Category, error) {
	return s.repo.StoreCategory(Category)
}

func (s *CategoryService) GetAllCategory() ([]dto.Category, error) {
	return s.repo.GetAllCategory()
}

func (s *CategoryService) ShowCategory(CategoryId int) (dto.Category, error) {
	return s.repo.ShowCategory(CategoryId)
}

func (s *CategoryService) DeleteCategory(CategoryId int) error {
	return s.repo.DeleteCategory(CategoryId)
}

func (s *CategoryService) UpdatedCategory(Category dto.CategoryUpdate, CategoryId int) error {
	return s.repo.UpdatedCategory(Category, CategoryId)
}
