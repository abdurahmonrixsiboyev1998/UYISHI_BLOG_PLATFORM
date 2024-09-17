package service

import (
	"blog/internal/auth"
	"blog/internal/models"
	"blog/internal/repository"
	"errors"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) Register(username, email, password string) error {
	hashPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Email: email,
		PasswordHash: hashPassword,
	}

	return s.Repo.CreateUser(&user)
}

func (s *UserService) Login(email, password string)(string, error) {
	user, err := s.Repo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := auth.CheckPassword(user.PasswordHash, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
