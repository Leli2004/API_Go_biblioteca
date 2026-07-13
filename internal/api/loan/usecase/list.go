package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type ListUC struct {
	repo loan.Repository
}

func NewListUC(repo loan.Repository) ListUC {
	return ListUC{repo: repo}
}

func (u *ListUC) Execute(input entity.LoanFilters) (error, entity.LoanList) {
	input.SetDefault()
	return u.repo.List(input)
}
