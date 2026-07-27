package repository

import (
	"context"

	"golang-for-devops/internal/models"
)

type LoanRepository interface {
	BorrowBook(
		context.Context,
		models.Loan,
	) error

	ReturnBook(
		context.Context,
		int,
	) error

	GetLoanByID(
		ctx context.Context,
		id int,
	) (models.Loan, error)

	GetLoans(
		ctx context.Context,
	) ([]models.Loan, error)
}
