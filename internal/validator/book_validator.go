package validator

import (
	"errors"
	"golang-for-devops/internal/models"
)

func ValidateBook(book models.Book) error {

	if book.Title == "" {
		return errors.New("title is required")
	}

	if book.Author == "" {
		return errors.New("author is required")
	}

	return nil
}
