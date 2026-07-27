package repository

import (
	"errors"

	"golang-for-devops/internal/auth"
	"golang-for-devops/internal/models"
)

type InMemoryUserRepository struct {
	users map[string]models.User
}

func NewUserRepository() *InMemoryUserRepository {
	hash, _ := auth.HashPassword("password123")
	return &InMemoryUserRepository{
		users: map[string]models.User{

			"admin": {
				Username:     "admin",
				PasswordHash: hash,
				Role:         auth.RoleAdmin,
			},

			"librarian": {
				Username:     "librarian",
				PasswordHash: hash,
				Role:         auth.RoleLibrarian,
			},

			"member": {
				Username:     "member",
				PasswordHash: hash,
				Role:         auth.RoleMember,
			},
		},
	}
}

func (r *InMemoryUserRepository) GetByUsername(username string) (*models.User, error) {

	user, ok := r.users[username]

	if !ok {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
