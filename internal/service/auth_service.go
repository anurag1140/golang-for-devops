package service

import (
	"errors"
	"golang-for-devops/internal/auth"
	"golang-for-devops/internal/models"
	"golang-for-devops/internal/repository"
	"time"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type AuthService struct {
	userRepo    repository.UserRepository
	refreshRepo repository.RefreshTokenRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
	refreshRepo repository.RefreshTokenRepository,
) *AuthService {

	return &AuthService{
		userRepo:    userRepo,
		refreshRepo: refreshRepo,
	}
}

func (s *AuthService) Login(username, password string) (*models.LoginResponse, error) {

	user, err := s.userRepo.GetByUsername(username)
	// var ErrInvalidCredentials = errors.New("invalid username or password")

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := auth.GenerateAccessToken(
		user.Username,
		user.Role,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken(
		user.Username,
		user.Role,
	)
	if err != nil {
		return nil, err
	}
	refresh := models.RefreshToken{
		Token:     refreshToken,
		Username:  user.Username,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		IsRevoked: false,
	}
	err = s.refreshRepo.Save(refresh)

	if err != nil {
		return nil, err
	}

	response := &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (s *AuthService) Refresh(
	refreshToken string,
) (*models.AccessTokenResponse, error) {

	// 1. Find refresh token in repository
	session, err := s.refreshRepo.Get(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// 2. Check if revoked
	if session.IsRevoked {
		return nil, errors.New("refresh token revoked")
	}

	// 3. Check expiration
	if time.Now().After(session.ExpiresAt) {
		return nil, errors.New("refresh token expired")
	}

	// 4. Validate JWT signature
	claims, err := auth.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// 5. Generate NEW access token
	accessToken, err := auth.GenerateAccessToken(
		claims.Username,
		claims.Role,
	)

	if err != nil {
		return nil, err
	}

	return &models.AccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}

func (s *AuthService) Logout(
	refreshToken string,
) error {

	return s.refreshRepo.Revoke(
		refreshToken,
	)
}
