package main

import (
	"golang-for-devops/internal/models"
	"golang-for-devops/internal/repository"
)

func seedBooks(repo repository.BookRepository) {

	repo.Add(models.Book{
		ID:        1,
		Title:     "Go in Action",
		Author:    "William",
		ISBN:      "12345",
		Available: true,
	})

	repo.Add(models.Book{
		ID:        2,
		Title:     "Microservices with Go",
		Author:    "John Smith",
		ISBN:      "68670",
		Available: true,
	})

	repo.Add(models.Book{
		ID:        3,
		Title:     "Some New Book",
		Author:    "Anurag",
		ISBN:      "23456",
		Available: false,
	})
}
