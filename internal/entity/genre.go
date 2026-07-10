package entity

type Genre struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description,omitempty" db:"description"`
	CreatedAt   *string `json:"created_at" db:"created_at"`
	UpdatedAt   *string `json:"updated_at" db:"updated_at"`
}

type GenreFilters struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (g *GenreFilters) SetDefault() {
	if g.Limit == 0 {
		g.Limit = 10
	}
}

type GenreList struct {
	Offset int      `json:"offset"`
	Limit  int      `json:"limit"`
	Data   []*Genre `json:"data"`
}
