package book_copie

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	List(ctx context.Context, tx *sqlx.Tx, input entity.BookCopyFilters) (context.Context, error, entity.BookCopyList)
	Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.BookCopy)
	Create(ctx context.Context, tx *sqlx.Tx, input entity.BookCopy) (context.Context, error, entity.BookCopy)
	Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.BookCopy) (context.Context, error, entity.BookCopy)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error)
}
