package publisher

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

//go:generate mockery --name=UseCase --dir=. --output=mocks --filename=mock_usecase.go --with-expecter=True

type UseCase interface {
	List(ctx context.Context, input entity.PublisherFilters) (context.Context, error, entity.PublisherList)
	Get(ctx context.Context, id int) (context.Context, error, entity.Publisher)
	Create(ctx context.Context, input entity.Publisher, claims *entity.AuthClaims) (context.Context, error, entity.Publisher)
	Update(ctx context.Context, id int, input entity.Publisher, claims *entity.AuthClaims) (context.Context, error, entity.Publisher)
	Delete(ctx context.Context, id int, claims *entity.AuthClaims) (context.Context, error)
}
