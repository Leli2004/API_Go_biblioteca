package genre

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UseCase interface {
	List(ctx context.Context, input entity.GenreFilters) (context.Context, error, entity.GenreList)
	Get(ctx context.Context, id int) (context.Context, error, entity.Genre)
	Create(ctx context.Context, input entity.Genre) (context.Context, error, entity.Genre)
	Update(ctx context.Context, id int, input entity.Genre) (context.Context, error, entity.Genre)
	Delete(ctx context.Context, id int) (context.Context, error)
}
