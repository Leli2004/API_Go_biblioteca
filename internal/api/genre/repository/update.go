package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct {
	db *sqlx.DB
}

func NewUpdateRepo(db *sqlx.DB) UpdateRepo {
	return UpdateRepo{db: db}
}

func (r *UpdateRepo) Execute(id int, input entity.Genre) (error, entity.Genre) {
	var genre entity.Genre
	err := r.db.Get(&genre, updateSql, input.Name, input.Description, id)
	if err != nil {
		return err, entity.Genre{}
	}

	return nil, genre
}

var updateSql = `
		UPDATE biblioteca.genres
		SET name = $1,
		    description = $2,
		    updated_at = NOW()
		WHERE id = $3
		RETURNING id, name, description, created_at, updated_at
	`
