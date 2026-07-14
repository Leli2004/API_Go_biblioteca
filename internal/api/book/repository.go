package book

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	List(ctx context.Context, tx *sqlx.Tx, input entity.BookFilters) (context.Context, error, entity.BookList)
	Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Book)
	Create(ctx context.Context, tx *sqlx.Tx, input entity.Book) (context.Context, error, entity.Book)
	Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Book) (context.Context, error, entity.Book)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error)
}
