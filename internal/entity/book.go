package entity

import "fmt"

type Book struct {
	Id              int         `json:"id" db:"id"`
	PublisherId     *int        `json:"publisher_id,omitempty" db:"publisher_id"`
	Title           string      `json:"title" db:"title"`
	PublicationYear *int        `json:"publication_year,omitempty" db:"publication_year"`
	Description     *string     `json:"description,omitempty" db:"description"`
	CreatedAt       *string     `json:"created_at" db:"created_at"`
	UpdatedAt       *string     `json:"updated_at" db:"updated_at"`
	AuthorIds       []int       `json:"author_ids,omitempty" db:"-"`
	GenreIds        []int       `json:"genre_ids,omitempty" db:"-"`
	Authors         []*Author   `json:"authors,omitempty" db:"-"`
	Genres          []*Genre    `json:"genres,omitempty" db:"-"`
	Copies          []*BookCopy `json:"copies,omitempty" db:"-"`
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

type BookCopy struct {
	Id        int     `json:"id" db:"id"`
	BookId    int     `json:"book_id" db:"book_id"`
	Barcode   string  `json:"barcode" db:"barcode"`
	Status    string  `json:"status" db:"status"` // available, loaned, maintenance, lost
	CreatedAt *string `json:"created_at" db:"created_at"`
	UpdatedAt *string `json:"updated_at" db:"updated_at"`
}

type BookList struct {
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
	Data   []*Book `json:"data"`
}
