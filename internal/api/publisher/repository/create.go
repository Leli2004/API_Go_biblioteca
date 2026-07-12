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

func (r *CreateRepo) Execute(input entity.Publisher) (error, entity.Publisher) {
	var publisher entity.Publisher
	err := r.db.Get(&publisher, createSql, input.Name, input.Website)
	if err != nil {
		return err, entity.Publisher{}
	}

	return nil, publisher
}

var createSql = `
	INSERT INTO biblioteca.publishers (name, website)
	VALUES ($1, $2)
	RETURNING id, name, website, created_at, updated_at
`
