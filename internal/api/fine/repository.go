package fine

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	ProcessOverdueLoans(ctx context.Context, tx *sqlx.Tx, amount float64, reason string) (context.Context, error, int)
}
