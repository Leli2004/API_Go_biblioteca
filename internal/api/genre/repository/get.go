package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct {
	db *sqlx.DB
}

func NewGetRepo(db *sqlx.DB) GetRepo {
	return GetRepo{db: db}
}

func (r *GetRepo) Execute(id int) (error, entity.Genre) {
	var genre entity.Genre
	err := r.db.Get(&genre, getSql, id)
	if err != nil {
		return err, entity.Genre{}
	}

	return nil, genre
}

var getSql = `
		SELECT
			id,
			name,
			description,
			created_at,
			updated_at
		FROM biblioteca.genres
		WHERE id = $1
	`
