package models

import "time"

type Loan struct {
	ID         int
	BookId     int
	MemberId   int
	IssueDate  time.Time
	ReturnDate time.Time
}
