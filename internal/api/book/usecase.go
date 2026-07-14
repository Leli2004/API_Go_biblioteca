package book

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UseCase interface {
	List(ctx context.Context, input entity.BookFilters) (context.Context, error, entity.BookList)
	Get(ctx context.Context, id int) (context.Context, error, entity.Book)
	Create(ctx context.Context, input entity.Book) (context.Context, error, entity.Book)
	Update(ctx context.Context, id int, input entity.Book) (context.Context, error, entity.Book)
	Delete(ctx context.Context, id int) (context.Context, error)
}
