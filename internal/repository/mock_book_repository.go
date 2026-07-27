package repository

import (
	"context"

	"golang-for-devops/internal/models"
)

type MockBookRepository struct {
	Books []models.Book
}

func NewMockBookRepository() *MockBookRepository {
	return &MockBookRepository{}
}

func (m *MockBookRepository) Add(
	ctx context.Context,
	book models.Book,
) error {

	m.Books = append(m.Books, book)
	return nil
}

func (m *MockBookRepository) GetAll(
	ctx context.Context,
) ([]models.Book, error) {

	return m.Books, nil
}

func (m *MockBookRepository) GetByID(
	ctx context.Context,
	id int,
) (*models.Book, error) {

	for _, b := range m.Books {
		if b.ID == id {
			return &b, nil
		}
	}

	return nil, nil
}

func (m *MockBookRepository) Update(
	ctx context.Context,
	book models.Book,
) error {

	return nil
}

func (m *MockBookRepository) Delete(
	ctx context.Context,
	id int,
) error {

	return nil
}

func (m *MockBookRepository) Search(
	ctx context.Context,
	query models.BookQuery,
) ([]models.Book, error) {

	return m.Books, nil
}
