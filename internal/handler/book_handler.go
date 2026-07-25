package handler

import (
	"encoding/json"
	"golang-for-devops/internal/models"
	"golang-for-devops/internal/service"
	"net/http"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {

	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {

	books := h.service.GetAllBooks()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) Books(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		h.GetBooks(w, r)

	case http.MethodPost:
		h.AddBook(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book

	// Convert JSON request into Go struct
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Call service
	if err := h.service.AddBook(r.Context(), book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// err = h.service.AddBook(book)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(book)
}
