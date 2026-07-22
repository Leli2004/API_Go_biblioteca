package reservation

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=Repository --dir=. --output=mocks --filename=mock_repository.go --with-expecter=True

type Repository interface {
	GetActiveByUserAndBook(ctx context.Context, tx *sqlx.Tx, userId, bookId int) (context.Context, error, entity.Reservation)
	GetNextPosition(ctx context.Context, tx *sqlx.Tx, bookId int) (context.Context, error, int)
	Create(ctx context.Context, tx *sqlx.Tx, input entity.Reservation) (context.Context, error, entity.Reservation)
}
