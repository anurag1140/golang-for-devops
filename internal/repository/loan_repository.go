package repository

import (
	"context"

	"golang-for-devops/internal/models"
)

type LoanRepository interface {
	IssueBook(
		context.Context,
		models.Loan,
	) error

	ReturnBook(
		context.Context,
		int,
	) error
}
