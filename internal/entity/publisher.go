package entity

type Publisher struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Website   *string `json:"website,omitempty"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}
