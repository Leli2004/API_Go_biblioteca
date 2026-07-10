package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct {
	db *sqlx.DB
}

func NewDeleteRepo(db *sqlx.DB) DeleteRepo {
	return DeleteRepo{
		db: db,
	}
}

func (r *DeleteRepo) Execute(id int) (error, entity.Author) {
	var author entity.Author
	err := r.db.Get(&author, deleteSql, id)
	if err != nil {
		return err, entity.Author{}
	}

	return nil, author
}

var deleteSql = `
		DELETE FROM biblioteca.authors
		WHERE id = $1
		RETURNING id, name, created_at, updated_at
	`
