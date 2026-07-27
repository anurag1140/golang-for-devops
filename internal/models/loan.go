package models

import "time"

type Loan struct {
	ID int

	BookID int

	MemberUsername string

	IssuedAt time.Time

	DueDate time.Time

	ReturnedAt *time.Time
}
