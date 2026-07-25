package service

import (
	"context"
	"encoding/json"
	"fmt"

	"golang-for-devops/internal/models"
	"golang-for-devops/internal/repository"
	"golang-for-devops/internal/storage"
	"golang-for-devops/internal/validator"
)

type BookService struct { // BookService struct
	repo    repository.BookRepository
	storage storage.Storage
}

func NewBookService(repo repository.BookRepository, storage storage.Storage) *BookService { //Constructor Injection.

	return &BookService{
		repo:    repo,
		storage: storage,
	}
}

func (s *BookService) AddBook(
	ctx context.Context,
	book models.Book,

) error { //dependency injection

	if err := validator.ValidateBook(book); err != nil {
		return err
	}

	s.repo.Add(book)

	data, err := json.Marshal(book)
	if err != nil {
		return err
	}

	key := fmt.Sprintf(
		"books/book-%d.json",
		book.ID,
	)

	err = s.storage.Upload(
		ctx,
		key,
		data,
	)

	if err != nil {
		return err
	}

	return nil

}

func (s *BookService) GetAllBooks() []models.Book {
	return s.repo.GetAll() //for now just delegation
}

func (s *BookService) GetBookByID(id int) (models.Book, error) {
	return s.repo.GetByID(id) //for now just delegation
}
