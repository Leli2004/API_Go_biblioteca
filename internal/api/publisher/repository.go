package publisher

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	List(ctx context.Context, tx *sqlx.Tx, input entity.PublisherFilters) (context.Context, error, entity.PublisherList)
	Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Publisher)
	Create(ctx context.Context, tx *sqlx.Tx, input entity.Publisher) (context.Context, error, entity.Publisher)
	Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Publisher) (context.Context, error, entity.Publisher)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error)
}
