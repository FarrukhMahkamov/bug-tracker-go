package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
)

const salt = "ERYncs%Cvbbd76SD%qadqw"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(user dto.User) (dto.UserForUi, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return s.repo.RegisterUser(user)
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
