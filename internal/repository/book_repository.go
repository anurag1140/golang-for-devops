package repository

import (
	"context"
	"golang-for-devops/internal/models"
)

type BookRepository interface {
	Add(context.Context, models.Book) error

	GetAll(context.Context) ([]models.Book, error)

	GetByID(context.Context, int) (*models.Book, error)

	Update(context.Context, models.Book) error

	Delete(context.Context, int) error

	Search(
		context.Context,
		models.BookQuery,
	) ([]models.Book, error)
}
