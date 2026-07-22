package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct{}

func NewListRepo() ListRepo {
	return ListRepo{}
}

func (r *ListRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.PublisherFilters) (context.Context, error, entity.PublisherList) {
	var publishers []*entity.Publisher
	err := tx.SelectContext(ctx, &publishers, listSql, input.Offset, input.Limit)
	if err != nil {
		return ctx, err, entity.PublisherList{}
	}

	return ctx, nil, entity.PublisherList{
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
