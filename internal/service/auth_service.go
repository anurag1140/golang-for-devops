package service

import (
	"errors"
	"golang-for-devops/internal/auth"
)

type AuthService struct {
}

func NewAuthService() *AuthService {

	return &AuthService{}
}

func (s *AuthService) Login(username, password string) (string, error) {

	if username != "admin" || password != "password123" {
		return "", errors.New("invalid username or password")
	}

	token, err := auth.GenerateToken(
		"admin",
		"librarian",
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
