package repository

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) RegisterUser(user dto.User) (dto.UserForUi, error) {
	var RegisteredUser dto.UserForUi

	row := r.db.QueryRow(query.RegisterUser, user.Name, user.UserName, user.Email, user.Password, user.Image)

	if err := row.Scan(
		&RegisteredUser.Id,
		&RegisteredUser.Name,
		&RegisteredUser.UserName,
		&RegisteredUser.Email,
		&RegisteredUser.Image,
	); err != nil {
		return dto.UserForUi{}, err
	}

	return RegisteredUser, nil
}
