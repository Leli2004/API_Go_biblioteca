package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type ListUC struct {
	repo user.Repository
}

func NewListUC(repo user.Repository) ListUC {
	return ListUC{repo: repo}
}

func (u *ListUC) Execute(input entity.UserFilters) (error, entity.UserList) {
	input.SetDefault()
	return u.repo.List(input)
}
