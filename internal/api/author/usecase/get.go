package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type GetUC struct {
	repo author.Repository
}

func NewGetUC(repo author.Repository) GetUC {
	return GetUC{repo: repo}
}

func (u *GetUC) Execute(id int) (error, entity.Author) {

	return u.repo.Get(id)
}
