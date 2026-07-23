package genre

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

//go:generate mockery --name=UseCase --dir=. --output=mocks --filename=mock_usecase.go --with-expecter=True

type UseCase interface {
	List(ctx context.Context, input entity.GenreFilters) (context.Context, error, entity.GenreList)
	Get(ctx context.Context, id int) (context.Context, error, entity.Genre)
	Create(ctx context.Context, input entity.Genre, claims *entity.AuthClaims) (context.Context, error, entity.Genre)
	Update(ctx context.Context, id int, input entity.Genre, claims *entity.AuthClaims) (context.Context, error, entity.Genre)
	Delete(ctx context.Context, id int, claims *entity.AuthClaims) (context.Context, error)
}
