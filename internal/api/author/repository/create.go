package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct {
	db *sqlx.DB
}

func NewCreateRepo(db *sqlx.DB) CreateRepo {
	return CreateRepo{
		db: db,
	}
}

func (r *CreateRepo) Execute(input entity.Author) (error, entity.Author) {
	var author entity.Author
	err := r.db.Get(&author, createSql, input.Name)
	if err != nil {
		return err, entity.Author{}
	}

	return nil, author
}

var createSql = `
		INSERT INTO biblioteca.authors (name)
		VALUES ($1)
		RETURNING id, name, created_at, updated_at
	`
