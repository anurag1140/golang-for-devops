package handler

import (
	"encoding/json"
	"golang-for-devops/internal/auth"
	"golang-for-devops/internal/models"
	"golang-for-devops/internal/service"
	"log"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {

	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) GetBooks(
	w http.ResponseWriter,
	r *http.Request,
) {

	books, err := h.service.GetAllBooks(
		r.Context(),
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Read authenticated user from context
	claims, ok := r.Context().
		Value(auth.UserContextKey).(*auth.Claims)

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) GetBookByID(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookByID(
		r.Context(),
		id,
	)

	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) UpdateBook(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	book.ID = id

	err = h.service.UpdateBook(
		r.Context(),
		book,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) DeleteBook(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteBook(
		r.Context(),
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
