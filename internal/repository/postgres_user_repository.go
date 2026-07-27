package repository

import (
	"context"
	"golang-for-devops/internal/models"

	"github.com/jackc/pgx/v5"
)

type PostgresUserRepository struct {
	db *pgx.Conn
}

func NewPostgresUserRepository(
	db *pgx.Conn,
) *PostgresUserRepository {

	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) GetByUsername(
	username string,
) (*models.User, error) {

	var user models.User

	err := r.db.QueryRow(
		context.Background(),
		`
		SELECT
			username,
			password_hash,
			role
		FROM users
		WHERE username=$1
		`,
		username,
	).Scan(
		&user.Username,
		&user.PasswordHash,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
