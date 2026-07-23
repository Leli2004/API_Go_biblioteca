package user

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

//go:generate mockery --name=UseCase --dir=. --output=mocks --filename=mock_usecase.go --with-expecter=True

type UseCase interface {
	List(ctx context.Context, input entity.UserFilters) (context.Context, error, entity.UserList)
	Get(ctx context.Context, id int) (context.Context, error, entity.User)
	Create(ctx context.Context, input entity.User, claims *entity.AuthClaims) (context.Context, error, entity.User)
	Update(ctx context.Context, id int, input entity.User, claims *entity.AuthClaims) (context.Context, error, entity.User)
	Delete(ctx context.Context, id int, claims *entity.AuthClaims) (context.Context, error)
}
