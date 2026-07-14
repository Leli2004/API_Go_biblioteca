package author

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	List(ctx context.Context, tx *sqlx.Tx, input entity.AuthorFilters) (context.Context, error, entity.AuthorList)
	Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Author)
	Create(ctx context.Context, tx *sqlx.Tx, input entity.Author) (context.Context, error, entity.Author)
	Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Author) (context.Context, error, entity.Author)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error)
}
