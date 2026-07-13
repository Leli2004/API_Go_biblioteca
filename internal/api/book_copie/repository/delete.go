package repository

import (
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct {
	db *sqlx.DB
}

func NewDeleteRepo(db *sqlx.DB) DeleteRepo {
	return DeleteRepo{db: db}
}

func (r *DeleteRepo) Execute(id int) (error, int) {
	var deletedId int
	err := r.db.Get(&deletedId, deleteSql, id)
	if err != nil {
		return err, 0
	}

	return nil, deletedId
}

var deleteSql = `
	DELETE FROM biblioteca.book_copies
	WHERE id = $1
	RETURNING id
`
