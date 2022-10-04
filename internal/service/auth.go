package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
	"github.com/form3tech-oss/jwt-go"
)

const salt = "ERYncs%Cvbbd76SD%qadqw"
const tokkenTTL = 12 * time.Hour
const signingKey = "396c0e76fec8eb5cb57e9d090e1aa3a20faf75a354b10c2b5f794551aaf47a4a"

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

type TokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func (s *AuthService) SignInUser(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, GeneratePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokkenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GetUser(email, password string) (dto.UserForUi, error) {
	user, err := s.repo.FindUser(email, GeneratePasswordHash(password))
	return user, err
}

func (s *AuthService) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type of *TokenClaims")
	}

	return claims.UserId, nil

}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
