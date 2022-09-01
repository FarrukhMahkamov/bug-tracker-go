package service

import "github.com/FarrukhMahkamov/bugtracker/internal/repository"

type Authortization interface {
}

type JobType interface {
}

type Service struct {
	Authortization
	JobType
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
