package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type GetUC struct {
	repo publisher.Repository
}

func NewGetUC(repo publisher.Repository) GetUC {
	return GetUC{repo: repo}
}

func (u *GetUC) Execute(id int) (error, entity.Publisher) {
	return u.repo.Get(id)
}
