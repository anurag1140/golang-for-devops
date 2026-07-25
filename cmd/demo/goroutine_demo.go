package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintBooks(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println("Book", i)
		time.Sleep(time.Second)
	}
}
