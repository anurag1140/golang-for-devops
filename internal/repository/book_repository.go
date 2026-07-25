package repository

import "golang-for-devops/internal/models"

type BookRepository interface {
	Add(book models.Book)

	GetAll() []models.Book

	GetByID(id int) (models.Book, error)
}
