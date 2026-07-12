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

func (r *UpdateRepo) Execute(id int, input entity.Publisher) (error, entity.Publisher) {
	var publisher entity.Publisher
	err := r.db.Get(&publisher, updateSql, input.Name, input.Website, id)
	if err != nil {
		return err, entity.Publisher{}
	}

	return nil, publisher
}

var updateSql = `
	UPDATE biblioteca.publishers
	SET name = $1,
	    website = $2,
	    updated_at = NOW()
	WHERE id = $3
	RETURNING id, name, website, created_at, updated_at
`
