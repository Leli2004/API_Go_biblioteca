package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type AuthorUC struct {
	listUC ListUC
	repo   author.Repository
}

func NewUseCase(repo author.Repository) AuthorUC {
	return AuthorUC{
		listUC: NewListUC(repo),
	}
}

func (u *AuthorUC) List(input entity.AuthorFilters) (error, entity.AuthorList) {
	return u.listUC.Execute(input)
}
