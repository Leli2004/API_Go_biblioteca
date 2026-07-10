package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct {
	db *sqlx.DB
}

func NewGetRepo(db *sqlx.DB) GetRepo {
	return GetRepo{
		db: db,
	}
}

func (r *GetRepo) Execute(id int) (error, entity.Author) {
	var author entity.Author
	err := r.db.Get(&author, getSql, id)
	if err != nil {
		return err, entity.Author{}
	}

	return nil, author
}

var getSql = `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM biblioteca.authors
		WHERE id = $1
	`
