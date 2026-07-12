package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct {
	db *sqlx.DB
}

func NewListRepo(db *sqlx.DB) ListRepo {
	return ListRepo{db: db}
}

func (r *ListRepo) Execute(input entity.PublisherFilters) (error, entity.PublisherList) {
	var publishers []*entity.Publisher
	err := r.db.Select(&publishers, listSql, input.Offset, input.Limit)
	if err != nil {
		return err, entity.PublisherList{}
	}

	return nil, entity.PublisherList{
		Offset: input.Offset,
		Limit:  helpers.GetMin(input.Limit, len(publishers)),
		Data:   publishers,
	}
}

var listSql = `
	SELECT
		id,
		name,
		website,
		created_at,
		updated_at
	FROM biblioteca.publishers
	ORDER BY id DESC
	OFFSET $1 LIMIT $2
`
