package loan

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type UseCase interface {
	Create(input entity.Loan) (error, entity.Loan)
	Return(loanId int, returnedAt *string) (error, entity.Loan)
	List(input entity.LoanFilters) (error, entity.LoanList)
	Get(id int) (error, entity.Loan)
	Delete(id int) (error, entity.Loan)
}
