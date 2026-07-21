package author

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

//go:generate mockery --name=UseCase --dir=. --output=mocks --filename=mock_usecase.go --with-expecter=True

type UseCase interface {
	List(ctx context.Context, input entity.AuthorFilters) (context.Context, error, entity.AuthorList)
	Get(ctx context.Context, id int) (context.Context, error, entity.Author)
	Create(ctx context.Context, input entity.Author) (context.Context, error, entity.Author)
	Update(ctx context.Context, id int, input entity.Author) (context.Context, error, entity.Author)
	Delete(ctx context.Context, id int) (context.Context, error)
}
