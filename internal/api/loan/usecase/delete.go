package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type DeleteUC struct {
	repo loan.Repository
}

func NewDeleteUC(repo loan.Repository) DeleteUC {
	return DeleteUC{repo: repo}
}

func (u *DeleteUC) Execute(id int) (error, entity.Loan) {
	return u.repo.Delete(id)
}
