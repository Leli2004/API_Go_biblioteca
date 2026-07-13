package loan

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type Repository interface {
	Create(input entity.Loan) (error, entity.Loan)
	Get(id int) (error, entity.Loan)
	Update(id int, input entity.Loan) (error, entity.Loan)
	GetActiveByUserAndBookCopy(userId int, bookCopyId int) (error, entity.Loan)
	List(input entity.LoanFilters) (error, entity.LoanList)
	Delete(id int) (error, entity.Loan)
}
