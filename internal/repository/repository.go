package repository

type Authortization interface {
}

type JobType interface {
}

type Repository struct {
	Authortization
	JobType
}

func NewRepository() *Repository {
	return &Repository{}
}
