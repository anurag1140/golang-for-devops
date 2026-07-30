package handler

import (
	"encoding/json"
	apierrors "golang-for-devops/internal/errors"
	"golang-for-devops/internal/models"
	"golang-for-devops/internal/service"
	"net/http"
	"strconv"
)

type LoanHandler struct {
	service *service.LoanService
}

func NewLoanHandler(
	service *service.LoanService,
) *LoanHandler {

	return &LoanHandler{
		service: service,
	}
}

func (h *LoanHandler) BorrowBook(
	w http.ResponseWriter,
	r *http.Request,
) {

	var loan models.Loan

	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {

		apierrors.WriteError(
			w,
			err,
		)

		return
	}

	err := h.service.BorrowBook(
		r.Context(),
		loan,
	)

	if err != nil {

		apierrors.WriteError(
			w,
			err,
		)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *LoanHandler) ReturnBook(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.Atoi(
		r.PathValue("bookId"),
	)

	if err != nil {

		apierrors.WriteError(
			w,
			err,
		)

		return
	}

	err = h.service.ReturnBook(
		r.Context(),
		id,
	)

	if err != nil {

		apierrors.WriteError(
			w,
			err,
		)

		return
	}

	w.WriteHeader(http.StatusOK)
}
