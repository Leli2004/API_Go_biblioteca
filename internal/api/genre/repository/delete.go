package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct {
	db *sqlx.DB
}

func NewDeleteRepo(db *sqlx.DB) DeleteRepo {
	return DeleteRepo{db: db}
}

func (r *DeleteRepo) Execute(id int) (error, entity.Genre) {
	var genre entity.Genre
	err := r.db.Get(&genre, deleteSql, id)
	if err != nil {
		return err, entity.Genre{}
	}

	return nil, genre
}

var deleteSql = `
		DELETE FROM biblioteca.genres
		WHERE id = $1
		RETURNING id, name, description, created_at, updated_at
	`
