package validator

import (
	"golang-for-devops/internal/models"

	apierrors "golang-for-devops/internal/errors"
)

func ValidateBook(book models.Book) error {

	if book.Title == "" {
		return apierrors.BadRequest(
			"title is required",
		)
	}

	if book.Author == "" {
		return apierrors.BadRequest(
			"author is required",
		)
	}

	return nil
}
