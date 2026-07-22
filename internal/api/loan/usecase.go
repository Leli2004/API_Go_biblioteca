package loan

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

//go:generate mockery --name=UseCase --dir=. --output=mocks --filename=mock_usecase.go --with-expecter=True

type UseCase interface {
	Create(ctx context.Context, input entity.Loan) (context.Context, error, entity.Loan)
	Return(ctx context.Context, loanId int, returnedAt *string) (context.Context, error, entity.Loan)
	List(ctx context.Context, input entity.LoanFilters) (context.Context, error, entity.LoanList)
	Get(ctx context.Context, id int) (context.Context, error, entity.Loan)
	Delete(ctx context.Context, id int) (context.Context, error, entity.Loan)
}
