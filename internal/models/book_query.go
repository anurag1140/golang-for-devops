package models

// type BookQuery struct {
// 	Page  int
// 	Size  int
// 	Title string
// 	Sort  string
// }

type BookQuery struct {
	Pagination

	Title     string
	Sort      string
	Author    string
	Available *bool
}
