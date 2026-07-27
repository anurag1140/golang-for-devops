package repository

import (
	"context"
	"errors"
	"golang-for-devops/internal/models"

	"github.com/jackc/pgx/v5"
)

type PostgresLoanRepository struct {
	db *pgx.Conn
}

func NewPostgresLoanRepository(
	db *pgx.Conn,
) *PostgresLoanRepository {

	return &PostgresLoanRepository{
		db: db,
	}
}

func (r *PostgresLoanRepository) BorrowBook(
	ctx context.Context,
	loan models.Loan,
) error {

	tx, err := r.db.Begin(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	//--------------------------------
	// Insert loan
	//--------------------------------

	_, err = tx.Exec(
		ctx,
		`
		INSERT INTO loans
		(book_id,
		member_username,
		due_date)
		VALUES($1,$2,$3)
		`,
		loan.BookID,
		loan.MemberUsername,
		loan.DueDate,
	)

	if err != nil {
		return err
	}

	//--------------------------------
	// Mark book unavailable
	//--------------------------------

	_, err = tx.Exec(
		ctx,
		`
		UPDATE books
		SET available=false
		WHERE id=$1
		`,
		loan.BookID,
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *PostgresLoanRepository) ReturnBook(
	ctx context.Context,
	bookID int,
) error {

	tx, err := r.db.Begin(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(
		ctx,
		`
		UPDATE loans
		SET returned_at=NOW()
		WHERE
			book_id=$1
			AND returned_at IS NULL
		`,
		bookID,
	)

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx,
		`
		UPDATE books
		SET available=true
		WHERE id=$1
		`,
		bookID,
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *PostgresLoanRepository) GetLoanByID(
	ctx context.Context,
	id int,
) (models.Loan, error) {

	return models.Loan{}, errors.New("not implemented")
}

func (r *PostgresLoanRepository) GetLoans(
	ctx context.Context,
) ([]models.Loan, error) {

	return nil, errors.New("not implemented")
}
