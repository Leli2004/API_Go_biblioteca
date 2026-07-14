package user

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	List(ctx context.Context, tx *sqlx.Tx, input entity.UserFilters) (context.Context, error, entity.UserList)
	Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.User)
	Create(ctx context.Context, tx *sqlx.Tx, input entity.User) (context.Context, error, entity.User)
	Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.User) (context.Context, error, entity.User)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error)
}
