package service

import (
	"context"
	"golang-for-devops/internal/models"
	"golang-for-devops/internal/repository"
)

type LoanService struct {
	repo repository.LoanRepository
}

func NewLoanService(
	repo repository.LoanRepository,
) *LoanService {

	return &LoanService{
		repo: repo,
	}
}

func (s *LoanService) IssueBook(
	ctx context.Context,
	loan models.Loan,
) error {

	return s.repo.IssueBook(
		ctx,
		loan,
	)
}

func (s *LoanService) ReturnBook(
	ctx context.Context,
	bookID int,
) error {

	return s.repo.ReturnBook(
		ctx,
		bookID,
	)
}
