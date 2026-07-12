package entity

import "fmt"

type Book struct {
	Id              int     `json:"id" db:"id"`
	PublisherId     *int    `json:"publisher_id,omitempty" db:"publisher_id"`
	Title           string  `json:"title" db:"title"`
	PublicationYear *int    `json:"publication_year,omitempty" db:"publication_year"`
	Description     *string `json:"description,omitempty" db:"description"`
	CreatedAt       *string `json:"created_at" db:"created_at"`
	UpdatedAt       *string `json:"updated_at" db:"updated_at"`
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return fmt.Errorf("Invalid field: Title is required")
	}
	return nil
}

type BookFilters struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (b *BookFilters) SetDefault() {
	if b.Limit == 0 {
		b.Limit = 10
	}
}

type BookList struct {
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
	Data   []*Book `json:"data"`
}
