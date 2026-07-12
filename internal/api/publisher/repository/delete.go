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

func (r *DeleteRepo) Execute(id int) (error, entity.Publisher) {
	var publisher entity.Publisher
	err := r.db.Get(&publisher, deleteSql, id)
	if err != nil {
		return err, entity.Publisher{}
	}

	return nil, publisher
}

var deleteSql = `
	DELETE FROM biblioteca.publishers
	WHERE id = $1
	RETURNING id, name, website, created_at, updated_at
`
