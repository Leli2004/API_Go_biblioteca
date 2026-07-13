package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type LoanUC struct {
	createUC CreateUC
	returnUC ReturnUC
	listUC   ListUC
	getUC    GetUC
	deleteUC DeleteUC
	repo     loan.Repository
}

func NewUseCase(repo loan.Repository) LoanUC {
	return LoanUC{
		createUC: NewCreateUC(repo),
		returnUC: NewReturnUC(repo),
		listUC:   NewListUC(repo),
		getUC:    NewGetUC(repo),
		deleteUC: NewDeleteUC(repo),
		repo:     repo,
	}
}

func (u *LoanUC) Create(input entity.Loan) (error, entity.Loan) {
	return u.createUC.Execute(input)
}

func (u *LoanUC) Return(loanId int, returnedAt *string) (error, entity.Loan) {
	return u.returnUC.Execute(loanId, returnedAt)
}

func (u *LoanUC) List(input entity.LoanFilters) (error, entity.LoanList) {
	return u.listUC.Execute(input)
}

func (u *LoanUC) Get(id int) (error, entity.Loan) {
	return u.repo.Get(id)
}

func (u *LoanUC) Delete(id int) (error, entity.Loan) {
	return u.repo.Delete(id)
}
