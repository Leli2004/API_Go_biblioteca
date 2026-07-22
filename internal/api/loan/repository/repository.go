package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type LoanRepo struct {
	create CreateRepo
	get    GetRepo
	update UpdateRepo
	active ActiveRepo
	list   ListRepo
	delete DeleteRepo
}

func NewRepository() *LoanRepo {
	return &LoanRepo{
		create: NewCreateRepo(),
		get:    NewGetRepo(),
		update: NewUpdateRepo(),
		active: NewActiveRepo(),
		list:   NewListRepo(),
		delete: NewDeleteRepo(),
	}
}

func (r *LoanRepo) Create(ctx context.Context, tx *sqlx.Tx, input entity.Loan) (context.Context, error, entity.Loan) {
	return r.create.Execute(ctx, tx, input)
}

func (r *LoanRepo) Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Loan) {
	return r.get.Execute(ctx, tx, id)
}

func (r *LoanRepo) Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Loan) (context.Context, error, entity.Loan) {
	return r.update.Execute(ctx, tx, id, input)
}

func (r *LoanRepo) GetActiveByUserAndBookCopy(ctx context.Context, tx *sqlx.Tx, userId int, bookCopyId int) (context.Context, error, entity.Loan) {
	return r.active.Execute(ctx, tx, userId, bookCopyId)
}

func (r *LoanRepo) List(ctx context.Context, tx *sqlx.Tx, input entity.LoanFilters) (context.Context, error, entity.LoanList) {
	return r.list.Execute(ctx, tx, input)
}

func (r *LoanRepo) Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Loan) {
	return r.delete.Execute(ctx, tx, id)
}
