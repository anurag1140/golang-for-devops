package main

import (
	"fmt"
	"golang-for-devops/internal/auth"
)

func main() {

	hash, _ := auth.HashPassword("password123")

	fmt.Println(hash)
}
