package handler

import (
	"encoding/json"
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

func (h *LoanHandler) IssueBook(
	w http.ResponseWriter,
	r *http.Request,
) {

	var loan models.Loan

	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {

		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)

		return
	}

	err := h.service.IssueBook(
		r.Context(),
		loan,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
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

		http.Error(
			w,
			"Invalid Book ID",
			http.StatusBadRequest,
		)

		return
	}

	err = h.service.ReturnBook(
		r.Context(),
		id,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(http.StatusOK)
}
