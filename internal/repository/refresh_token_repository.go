package repository

import "golang-for-devops/internal/models"

type RefreshTokenRepository interface {
	Save(token models.RefreshToken) error

	Get(token string) (*models.RefreshToken, error)

	Revoke(token string) error

	DeleteExpired() error
}
