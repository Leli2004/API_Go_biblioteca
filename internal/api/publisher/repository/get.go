package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct {
	db *sqlx.DB
}

func NewGetRepo(db *sqlx.DB) GetRepo {
	return GetRepo{db: db}
}

func (r *GetRepo) Execute(id int) (error, entity.Publisher) {
	var publisher entity.Publisher
	err := r.db.Get(&publisher, getSql, id)
	if err != nil {
		return err, entity.Publisher{}
	}

	return nil, publisher
}

var getSql = `
	SELECT
		id,
		name,
		website,
		created_at,
		updated_at
	FROM biblioteca.publishers
	WHERE id = $1
`
