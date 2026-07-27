package repository

import (
	"errors"
	"golang-for-devops/internal/models"
)

type InMemoryRefreshTokenRepository struct {
	tokens map[string]models.RefreshToken
}

// constructor
func NewRefreshTokenRepository() *InMemoryRefreshTokenRepository {

	return &InMemoryRefreshTokenRepository{
		tokens: make(map[string]models.RefreshToken),
	}
}

// save
// func (r *InMemoryRefreshTokenRepository) Save(
// 	token models.RefreshToken,
// ) error {

// 	r.tokens[token.Token] = token

// 	return nil
// }

// Implement get
func (r *InMemoryRefreshTokenRepository) Get(
	token string,
) (*models.RefreshToken, error) {

	refreshToken, ok := r.tokens[token]

	if !ok {
		return nil, errors.New("refresh token not found")
	}

	return &refreshToken, nil
}

func (r *InMemoryRefreshTokenRepository) Revoke(
	token string,
) error {

	storedToken, ok := r.tokens[token]

	if !ok {
		return errors.New("token not found")
	}

	storedToken.IsRevoked = true

	r.tokens[token] = storedToken

	return nil
}
