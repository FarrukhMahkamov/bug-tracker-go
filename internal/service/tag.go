package service

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

type TagService struct {
	repo repository.Tag
}

func NewTagService(repo repository.Tag) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) StoreTag(Tag dto.Tag) (dto.Tag, error) {
	return s.repo.StoreTag(Tag)
}

func (s *TagService) GetAllTag() ([]dto.Tag, error) {
	return s.repo.GetAllTag()
}

func (s *TagService) ShowTag(TagId int) (dto.Tag, error) {
	return s.repo.ShowTag(TagId)
}

func (s *TagService) DeleteTag(TagId int) error {
	return s.repo.DeleteTag(TagId)
}

func (s *TagService) UpdatedTag(Tag dto.TagUpdate, TagId int) error {
	return s.repo.UpdatedTag(Tag, TagId)
}
