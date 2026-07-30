package handler

import (
	"encoding/json"
	"golang-for-devops/internal/auth"
	"golang-for-devops/internal/models"
	"golang-for-devops/internal/service"
	"log"
	"net/http"
	"strconv"

	apierrors "golang-for-devops/internal/errors"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {

	return &BookHandler{
		service: service,
	}
}

// GetBooks godoc
//
// @Summary Get books
// @Description Returns all books
// @Tags Books
// @Produce json
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Param title query string false "Filter by title"
// @Param author query string false "Filter by author"
// @Success 200 {array} models.Book
// @Failure 500 {object} models.ErrorResponse
// @Router /books [get]
func (h *BookHandler) GetBooks(
	w http.ResponseWriter,
	r *http.Request,
) {

	query := ParseBookQuery(r)

	books, err := h.service.SearchBooks(
		r.Context(),
		query,
	)

	if err != nil {

		apierrors.WriteError(
			w,
			err,
		)

		return
	}

	apierrors.WriteJSON(
		w,
		http.StatusCreated,
		books,
	)
}

// AddBook godoc
//
// @Summary Add book
// @Description Adds a new book
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.Book true "Book"
// @Success 201 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /books [post]
func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	// Read authenticated user from context
	claims, ok := r.Context().
		Value(auth.UserContextKey).(*auth.Claims)

	if !ok {
		apierrors.WriteError(
			w,
			apierrors.Unauthorized(apierrors.CodeUnauthorized),
		)
		return
	}

	// Temporary logging
	log.Printf(
		"User %s (%s) is adding a book",
		claims.Username,
		claims.Role,
	)

	// Business logic
	if err := h.service.AddBook(r.Context(), book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	apierrors.WriteJSON(
		w,
		http.StatusCreated,
		book,
	)
}

// GetBookByID godoc
//
// @Summary Get book by ID
// @Description Returns one book
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	book, err := h.service.GetBookByID(
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

	apierrors.WriteJSON(
		w,
		http.StatusCreated,
		book,
	)
}

func (h *BookHandler) UpdateBook(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	book.ID = id

	err = h.service.UpdateBook(
		r.Context(),
		book,
	)

	if err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	apierrors.WriteJSON(
		w,
		http.StatusCreated,
		book,
	)
}

func (h *BookHandler) DeleteBook(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	err = h.service.DeleteBook(
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

	w.WriteHeader(http.StatusNoContent)
}
