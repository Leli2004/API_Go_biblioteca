package book_copie

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

//go:generate mockery --name=UseCase --dir=. --output=mocks --filename=mock_usecase.go --with-expecter=True

type UseCase interface {
	List(ctx context.Context, input entity.BookCopyFilters) (context.Context, error, entity.BookCopyList)
	Get(ctx context.Context, id int) (context.Context, error, entity.BookCopy)
	Create(ctx context.Context, input entity.BookCopy) (context.Context, error, entity.BookCopy)
	Update(ctx context.Context, id int, input entity.BookCopy) (context.Context, error, entity.BookCopy)
	Delete(ctx context.Context, id int) (context.Context, error)
}
