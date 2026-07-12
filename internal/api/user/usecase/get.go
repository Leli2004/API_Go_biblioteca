package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type GetUC struct {
	repo user.Repository
}

func NewGetUC(repo user.Repository) GetUC {
	return GetUC{repo: repo}
}

func (u *GetUC) Execute(id int) (error, entity.User) {
	return u.repo.Get(id)
}
