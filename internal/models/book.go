package models

import "fmt"

//	type Book struct {
//		ID        int
//		Title     string
//		Author    string
//		ISBN      string
//		Available bool
//	}
type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Available bool   `json:"available"`
}

func (b Book) Print() {

	fmt.Println("Book : ", b.Title)

}
