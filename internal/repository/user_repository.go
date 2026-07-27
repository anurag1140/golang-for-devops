package repository

import "golang-for-devops/internal/models"

type UserRepository interface {
	GetByUsername(username string) (*models.User, error)
}
