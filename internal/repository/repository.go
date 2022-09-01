package repository

import "github.com/jmoiron/sqlx"

type Authortization interface {
}

type JobType interface {
}

type Repository struct {
	Authortization
	JobType
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
