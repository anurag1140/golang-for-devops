package repository

import (
	"errors"
	"golang-for-devops/internal/models"
	"sync"
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
func (r *InMemoryBookRepository) Add(book models.Book) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.books = append(r.books, book) // append a new slice header and always returns back
}
func (r *InMemoryBookRepository) GetAll() []models.Book {

	return r.books
}

func (r *InMemoryBookRepository) GetByID(id int) (models.Book, error) {

	for _, book := range r.books { //foreach - range is used . _ means - I do not need index.

		if book.ID == id {
			return book, nil
		}
	}

	return models.Book{}, errors.New("book not found.") // error handling - Go returns an empty struct plus an error
}
