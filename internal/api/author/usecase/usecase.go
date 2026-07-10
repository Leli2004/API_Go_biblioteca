package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type AuthorUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     author.Repository
}

func NewUseCase(repo author.Repository) AuthorUC {
	return AuthorUC{
		listUC:   NewListUC(repo),
		getUC:    NewGetUC(repo),
		createUC: NewCreateUC(repo),
		updateUC: NewUpdateUC(repo),
		deleteUC: NewDeleteUC(repo),
		repo:     repo,
	}
}

func (u *AuthorUC) List(input entity.AuthorFilters) (error, entity.AuthorList) {
	return u.listUC.Execute(input)
}

func (u *AuthorUC) Get(id int) (error, entity.Author) {
	return u.getUC.Execute(id)
}

func (u *AuthorUC) Create(input entity.Author) (error, entity.Author) {
	return u.createUC.Execute(input)
}

func (u *AuthorUC) Update(id int, input entity.Author) (error, entity.Author) {
	return u.updateUC.Execute(id, input)
}

func (u *AuthorUC) Delete(id int) error {
	return u.deleteUC.Execute(id)
}
