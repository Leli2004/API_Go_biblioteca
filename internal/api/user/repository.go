package user

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=Repository --dir=. --output=mocks --filename=mock_repository.go --with-expecter=True

type Repository interface {
	List(ctx context.Context, tx *sqlx.Tx, input entity.UserFilters) (context.Context, error, entity.UserList)
	Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.User)
	Create(ctx context.Context, tx *sqlx.Tx, input entity.User) (context.Context, error, entity.User)
	Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.User) (context.Context, error, entity.User)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error)
}
