package entity

import (
	"fmt"
	"time"
)

type Loan struct {
	Id         int     `json:"id" db:"id"`
	UserId     int     `json:"user_id" db:"user_id"`
	BookCopyId int     `json:"book_copy_id" db:"book_copy_id"`
	LoanDate   *string `json:"loan_date,omitempty" db:"loan_date"`
	DueDate    *string `json:"due_date,omitempty" db:"due_date"`
	ReturnedAt *string `json:"returned_at,omitempty" db:"returned_at"`
	Status     string  `json:"status" db:"status"` // active, returned
	CreatedAt  *string `json:"created_at" db:"created_at"`
	UpdatedAt  *string `json:"updated_at" db:"updated_at"`
}

func (l *Loan) Validate() error {
	if l.UserId <= 0 {
		return fmt.Errorf("Invalid field: UserId is required")
	}
	if l.BookCopyId <= 0 {
		return fmt.Errorf("Invalid field: BookCopyId is required")
	}
	if l.DueDate == nil || *l.DueDate == "" {
		return fmt.Errorf("Invalid field: DueDate is required")
	}
	return nil
}

func (l *Loan) SetDefault() {
	if l.Status == "" {
		l.Status = "active"
	}
	if l.LoanDate == nil || *l.LoanDate == "" {
		t := time.Now().UTC().Format(time.RFC3339)
		l.LoanDate = &t
	}
}

type LoanFilters struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (f *LoanFilters) SetDefault() {
	if f.Limit == 0 {
		f.Limit = 10
	}
}

type LoanList struct {
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
	Data   []*Loan `json:"data"`
}
