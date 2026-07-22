package entity

import "fmt"

type Publisher struct {
	Id        int     `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Website   *string `json:"website,omitempty" db:"website"`
	CreatedAt *string `json:"created_at" db:"created_at"`
	UpdatedAt *string `json:"updated_at" db:"updated_at"`
}

func (p *Publisher) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("Invalid field: Name is required")
	}
	return nil
}

type PublisherFilters struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (p *PublisherFilters) SetDefault() {
	if p.Limit == 0 {
		p.Limit = 10
	}
}

type PublisherList struct {
	Offset int          `json:"offset"`
	Limit  int          `json:"limit"`
	Data   []*Publisher `json:"data"`
}
