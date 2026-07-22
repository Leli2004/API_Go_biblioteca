package entity

import "fmt"

type BookCopyFilters struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (b *BookCopyFilters) SetDefault() {
	if b.Limit == 0 {
		b.Limit = 10
	}
}

type BookCopyList struct {
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Data   []*BookCopy `json:"data"`
}

func (b *BookCopy) Validate() error {
	if b.BookId <= 0 {
		return fmt.Errorf("Invalid field: BookId is required")
	}
	if b.Barcode == "" {
		return fmt.Errorf("Invalid field: Barcode is required")
	}
	return nil
}
