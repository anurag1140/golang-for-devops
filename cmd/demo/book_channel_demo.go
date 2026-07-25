package main

import (
	"fmt"
	"golang-for-devops/internal/models"
	"time"
)

func GenerateBook(ch chan models.Book) {

	fmt.Println("Generating Book...")

	time.Sleep(2 * time.Second)

	book := models.Book{
		ID:        10,
		Title:     "Go Concurrency",
		Author:    "Priyanka",
		ISBN:      "77777",
		Available: true,
	}

	ch <- book
}
