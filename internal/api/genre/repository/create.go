package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct {
	db *sqlx.DB
}

func NewCreateRepo(db *sqlx.DB) CreateRepo {
	return CreateRepo{db: db}
}

func (r *CreateRepo) Execute(input entity.Genre) (error, entity.Genre) {
	var genre entity.Genre
	err := r.db.Get(&genre, createSql, input.Name, input.Description)
	if err != nil {
		return err, entity.Genre{}
	}

	return nil, genre
}

var createSql = `
		INSERT INTO biblioteca.genres (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description, created_at, updated_at
	`
