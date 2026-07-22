package genre

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=Repository --dir=. --output=mocks --filename=mock_repository.go --with-expecter=True

type Repository interface {
	List(ctx context.Context, tx *sqlx.Tx, input entity.GenreFilters) (context.Context, error, entity.GenreList)
	Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Genre)
	Create(ctx context.Context, tx *sqlx.Tx, input entity.Genre) (context.Context, error, entity.Genre)
	Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Genre) (context.Context, error, entity.Genre)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error)
}
