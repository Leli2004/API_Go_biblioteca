package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type GetUC struct {
	repo loan.Repository
}

func NewGetUC(repo loan.Repository) GetUC {
	return GetUC{repo: repo}
}

func (u *GetUC) Execute(id int) (error, entity.Loan) {
	return u.repo.Get(id)
}
