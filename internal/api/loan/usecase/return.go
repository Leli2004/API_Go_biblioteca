package usecase

import (
	"time"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type ReturnUC struct {
	repo loan.Repository
}

func NewReturnUC(repo loan.Repository) ReturnUC {
	return ReturnUC{repo: repo}
}

func (u *ReturnUC) Execute(loanId int, returnedAt *string) (error, entity.Loan) {
	err, ln := u.repo.Get(loanId)
	if err != nil {
		return err, entity.Loan{}
	}

	if returnedAt == nil || *returnedAt == "" {
		t := time.Now().UTC().Format(time.RFC3339)
		returnedAt = &t
	}

	ln.ReturnedAt = returnedAt
	ln.Status = "returned"

	err, updated := u.repo.Update(loanId, ln)
	if err != nil {
		return err, entity.Loan{}
	}

	return nil, updated
}
