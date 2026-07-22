package fine

import "context"

type UseCase interface {
	ProcessOverdueLoans(ctx context.Context) (context.Context, error, int)
}
