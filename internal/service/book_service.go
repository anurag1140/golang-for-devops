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

	if err := s.repo.Add(ctx, book); err != nil {
		return err
	}

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

func (s *BookService) GetAllBooks(
	ctx context.Context,
) ([]models.Book, error) {

	return s.repo.GetAll(ctx)
}

func (s *BookService) GetBookByID(
	ctx context.Context,
	id int,
) (*models.Book, error) {

	return s.repo.GetByID(ctx, id)
}

func (s *BookService) UpdateBook(
	ctx context.Context,
	book models.Book,
) error {

	if err := validator.ValidateBook(book); err != nil {
		return err
	}

	return s.repo.Update(ctx, book)
}

func (s *BookService) DeleteBook(
	ctx context.Context,
	id int,
) error {

	return s.repo.Delete(ctx, id)
}
