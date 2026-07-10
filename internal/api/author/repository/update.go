package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct {
	db *sqlx.DB
}

func NewUpdateRepo(db *sqlx.DB) UpdateRepo {
	return UpdateRepo{
		db: db,
	}
}

func (r *UpdateRepo) Execute(id int, input entity.Author) (error, entity.Author) {
	var author entity.Author
	err := r.db.Get(&author, updateSql, input.Name, id)
	if err != nil {
		return err, entity.Author{}
	}

	return nil, author
}

var updateSql = `
		UPDATE biblioteca.authors
		SET name = $1,
		updated_at = NOW()
		WHERE id = $2
		RETURNING id, name, created_at, updated_at
	`
