package repository

import (
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

func NewRepository(db *sqlx.DB) *LoanRepo {
	return &LoanRepo{
		create: NewCreateRepo(db),
		get:    NewGetRepo(db),
		update: NewUpdateRepo(db),
		active: NewActiveRepo(db),
		list:   NewListRepo(db),
		delete: NewDeleteRepo(db),
	}
}

func (r *LoanRepo) Create(input entity.Loan) (error, entity.Loan) {
	return r.create.Execute(input)
}

func (r *LoanRepo) Get(id int) (error, entity.Loan) {
	return r.get.Execute(id)
}

func (r *LoanRepo) Update(id int, input entity.Loan) (error, entity.Loan) {
	return r.update.Execute(id, input)
}

func (r *LoanRepo) GetActiveByUserAndBookCopy(userId int, bookCopyId int) (error, entity.Loan) {
	return r.active.Execute(userId, bookCopyId)
}

func (r *LoanRepo) List(input entity.LoanFilters) (error, entity.LoanList) {
	return r.list.Execute(input)
}

func (r *LoanRepo) Delete(id int) (error, entity.Loan) {
	return r.delete.Execute(id)
}
