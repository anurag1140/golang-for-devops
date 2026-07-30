package repository

import (
	"context"
	"errors"
	"golang-for-devops/internal/models"
	"sync"

	apierrors "golang-for-devops/internal/errors"
)

type InMemoryBookRepository struct { //InMemoryBookRepository structure
	books []models.Book // field - slice/dynamic array (a list) - books can have many books

	mu sync.Mutex
}

// Factory function -- similar to our constructor
func NewBookRepository() *InMemoryBookRepository { //This method belongs to the repository

	return &InMemoryBookRepository{ // Create a repository & Return its address.The & gives the address.

		books: []models.Book{},
	}
}
func (r *InMemoryBookRepository) Add(
	ctx context.Context,
	book models.Book,
) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.books = append(r.books, book)

	return nil
}

func (r *InMemoryBookRepository) GetAll(
	ctx context.Context,
) ([]models.Book, error) {

	return r.books, nil
}

func (r *InMemoryBookRepository) GetByID(
	ctx context.Context,
	id int,
) (*models.Book, error) {

	for _, book := range r.books {

		if book.ID == id {

			return &book, nil
		}
	}

	return nil, errors.New("book not found")
}

func (r *InMemoryBookRepository) Update(
	ctx context.Context,
	updated models.Book,
) error {

	for i := range r.books {

		if r.books[i].ID == updated.ID {

			r.books[i] = updated

			return nil
		}
	}

	return apierrors.BadRequest("Book not found")
}

func (r *InMemoryBookRepository) Delete(
	ctx context.Context,
	id int,
) error {

	for i := range r.books {

		if r.books[i].ID == id {

			r.books = append(
				r.books[:i],
				r.books[i+1:]...,
			)

			return nil
		}
	}

	return apierrors.BadRequest(
		"Book not found",
	)
}

func (r *InMemoryBookRepository) Search(
	ctx context.Context,
	query models.BookQuery,
) ([]models.Book, error) {

	// For now just return all books.
	// We'll enhance filtering later if needed.

	return r.GetAll(ctx)
}
