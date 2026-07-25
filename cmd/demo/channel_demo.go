package main

func SendMessage(ch chan string) {

	ch <- "Hello from Goroutine"

}
