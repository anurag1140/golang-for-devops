package repository

import (
	"context"

	"golang-for-devops/internal/models"

	"github.com/jackc/pgx/v5"
)

type PostgresBookRepository struct {
	db *pgx.Conn
}

func NewPostgresBookRepository(
	db *pgx.Conn,
) *PostgresBookRepository {

	return &PostgresBookRepository{
		db: db,
	}
}

func (r *PostgresBookRepository) Add(
	ctx context.Context,
	book models.Book,
) error {

	_, err := r.db.Exec(
		ctx,
		`
        INSERT INTO books
        (id,title,author,isbn,available)
        VALUES($1,$2,$3,$4,$5)
        `,
		book.ID,
		book.Title,
		book.Author,
		book.ISBN,
		book.Available,
	)

	return err
}

func (r *PostgresBookRepository) GetAll(
	ctx context.Context,
) ([]models.Book, error) {

	rows, err := r.db.Query(
		ctx,
		`
		SELECT
			id,
			title,
			author,
			isbn,
			available
		FROM books
		ORDER BY id
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []models.Book

	for rows.Next() {

		var book models.Book

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.ISBN,
			&book.Available,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *PostgresBookRepository) GetByID(
	ctx context.Context,
	id int,
) (*models.Book, error) {

	var book models.Book

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			title,
			author,
			isbn,
			available
		FROM books
		WHERE id=$1
		`,
		id,
	).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.ISBN,
		&book.Available,
	)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *PostgresBookRepository) Update(
	ctx context.Context,
	book models.Book,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		UPDATE books
		SET
			title=$1,
			author=$2,
			isbn=$3,
			available=$4
		WHERE id=$5
		`,
		book.Title,
		book.Author,
		book.ISBN,
		book.Available,
		book.ID,
	)

	return err
}

func (r *PostgresBookRepository) Delete(
	ctx context.Context,
	id int,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		DELETE FROM books
		WHERE id=$1
		`,
		id,
	)

	return err
}
