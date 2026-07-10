package entity

import "fmt"

type Author struct {
	Id        int     `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	CreatedAt *string `json:"created_at" db:"created_at"`
	UpdatedAt *string `json:"updated_at" db:"updated_at"`
}

func (a *Author) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("Invalid field: Name is required")
	}
	return nil
}

type AuthorFilters struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (a *AuthorFilters) SetDefault() {
	if a.Limit == 0 {
		a.Limit = 10
	}
}

type AuthorList struct {
	Offset int       `json:"offset"`
	Limit  int       `json:"limit"`
	Data   []*Author `json:"data"`
}
