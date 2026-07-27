package repository

import (
	"context"
	"fmt"
	"strconv"

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

func (r *PostgresBookRepository) Search(
	ctx context.Context,
	query models.BookQuery,
) ([]models.Book, error) {

	sql := `
	SELECT
		id,
		title,
		author,
		isbn,
		available
	FROM books
	WHERE 1=1
	`

	args := []any{}
	index := 1

	// ---------- Filters ----------

	if query.Title != "" {

		sql += " AND LOWER(title) LIKE LOWER($" +
			strconv.Itoa(index) + ")"

		args = append(
			args,
			"%"+query.Title+"%",
		)

		index++
	}

	if query.Author != "" {

		sql += " AND LOWER(author) LIKE LOWER($" +
			strconv.Itoa(index) + ")"

		args = append(
			args,
			"%"+query.Author+"%",
		)

		index++
	}

	// ---------- Sorting ----------

	allowedSort := map[string]string{
		"id":     "id",
		"title":  "title",
		"author": "author",
	}

	sortColumn := "id"

	if value, ok := allowedSort[query.Sort]; ok {
		sortColumn = value
	}

	sql += " ORDER BY " + sortColumn

	// ---------- Pagination ----------

	if query.Page < 1 {
		query.Page = 1
	}

	if query.Size < 1 {
		query.Size = 10
	}

	if query.Size > 100 {
		query.Size = 100
	}

	offset := (query.Page - 1) * query.Size

	sql += fmt.Sprintf(
		" LIMIT $%d OFFSET $%d",
		index,
		index+1,
	)

	args = append(
		args,
		query.Size,
		offset,
	)

	// ---------- Execute ----------

	rows, err := r.db.Query(
		ctx,
		sql,
		args...,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []models.Book

	for rows.Next() {

		var book models.Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.ISBN,
			&book.Available,
		); err != nil {
			return nil, err
		}

		books = append(
			books,
			book,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
