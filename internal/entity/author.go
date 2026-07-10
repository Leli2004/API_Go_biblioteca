package entity

type Author struct {
	Id        int     `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	CreatedAt *string `json:"created_at" db:"created_at"`
	UpdatedAt *string `json:"updated_at" db:"updated_at"`
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
