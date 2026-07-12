package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type GetUC struct {
	repo book.Repository
}

func NewGetUC(repo book.Repository) GetUC {
	return GetUC{repo: repo}
}

func (u *GetUC) Execute(id int) (error, entity.Book) {
	return u.repo.Get(id)
}
