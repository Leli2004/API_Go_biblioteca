package loan

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=Repository --dir=. --output=mocks --filename=mock_repository.go --with-expecter=True

type Repository interface {
	Create(ctx context.Context, tx *sqlx.Tx, input entity.Loan) (context.Context, error, entity.Loan)
	Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Loan)
	Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Loan) (context.Context, error, entity.Loan)
	GetActiveByUserAndBookCopy(ctx context.Context, tx *sqlx.Tx, userId int, bookCopyId int) (context.Context, error, entity.Loan)
	List(ctx context.Context, tx *sqlx.Tx, input entity.LoanFilters) (context.Context, error, entity.LoanList)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Loan)
}
