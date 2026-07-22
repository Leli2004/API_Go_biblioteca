package usecase

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type LoanUC struct {
	createUC CreateUC
	returnUC ReturnUC
	listUC   ListUC
	getUC    GetUC
	deleteUC DeleteUC
	repo     loan.Repository
}

func NewUseCase(db *sqlx.DB, repo loan.Repository) *LoanUC {
	return &LoanUC{
		createUC: NewCreateUC(db, repo),
		returnUC: NewReturnUC(db, repo),
		listUC:   NewListUC(db, repo),
		getUC:    NewGetUC(db, repo),
		deleteUC: NewDeleteUC(db, repo),
		repo:     repo,
	}
}

func (u *LoanUC) Create(ctx context.Context, input entity.Loan) (context.Context, error, entity.Loan) {
	return u.createUC.Execute(ctx, input)
}

func (u *LoanUC) Return(ctx context.Context, loanId int, returnedAt *string) (context.Context, error, entity.Loan) {
	return u.returnUC.Execute(ctx, loanId, returnedAt)
}

func (u *LoanUC) List(ctx context.Context, input entity.LoanFilters) (context.Context, error, entity.LoanList) {
	return u.listUC.Execute(ctx, input)
}

func (u *LoanUC) Get(ctx context.Context, id int) (context.Context, error, entity.Loan) {
	return u.getUC.Execute(ctx, id)
}

func (u *LoanUC) Delete(ctx context.Context, id int) (context.Context, error, entity.Loan) {
	return u.deleteUC.Execute(ctx, id)
}
