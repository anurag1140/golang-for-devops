package repository

import (
	"context"
	"errors"
	"golang-for-devops/internal/models"

	"github.com/jackc/pgx/v5"
)

type PostgresRefreshTokenRepository struct {
	db *pgx.Conn
}

func NewPostgresRefreshTokenRepository(
	db *pgx.Conn,
) *PostgresRefreshTokenRepository {

	return &PostgresRefreshTokenRepository{
		db: db,
	}
}

func (r *PostgresRefreshTokenRepository) Save(
	token models.RefreshToken,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`
		INSERT INTO refresh_tokens
		(token, username, expires_at, revoked)
		VALUES ($1,$2,$3,$4)
		`,
		token.Token,
		token.Username,
		token.ExpiresAt,
		token.IsRevoked,
	)

	return err
}

func (r *PostgresRefreshTokenRepository) Get(
	token string,
) (*models.RefreshToken, error) {

	var refreshToken models.RefreshToken

	err := r.db.QueryRow(
		context.Background(),
		`
		SELECT
			token,
			username,
			expires_at,
			revoked
		FROM refresh_tokens
		WHERE token=$1
		`,
		token,
	).Scan(
		&refreshToken.Token,
		&refreshToken.Username,
		&refreshToken.ExpiresAt,
		&refreshToken.IsRevoked,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("refresh token not found")
	}

	return &refreshToken, nil
}

func (r *PostgresRefreshTokenRepository) Revoke(
	token string,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`
		UPDATE refresh_tokens
		SET revoked = true
		WHERE token = $1
		`,
		token,
	)

	return err
}

func (r *PostgresRefreshTokenRepository) DeleteExpired() error {

	_, err := r.db.Exec(
		context.Background(),
		`
		DELETE FROM refresh_tokens
		WHERE expires_at < NOW()
		`,
	)

	return err
}
