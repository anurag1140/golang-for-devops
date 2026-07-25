package main

import (
	"fmt"
	"golang-for-devops/internal/models"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go PrintBooks(&wg)

	fmt.Println("Server started")

	wg.Wait()

	fmt.Println("All goroutines finished")

	ch := make(chan string)

	go SendMessage(ch)

	message := <-ch

	fmt.Println(message)

	bookChannel := make(chan models.Book)
	go GenerateBook(bookChannel)

	fmt.Println("Waiting for Book...")

	book := <-bookChannel

	fmt.Println("Received Book")

	fmt.Println(book)

}
