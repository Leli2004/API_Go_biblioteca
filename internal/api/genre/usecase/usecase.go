package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type GenreUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     genre.Repository
}

func NewUseCase(repo genre.Repository) GenreUC {
	return GenreUC{
		listUC:   NewListUC(repo),
		getUC:    NewGetUC(repo),
		createUC: NewCreateUC(repo),
		updateUC: NewUpdateUC(repo),
		deleteUC: NewDeleteUC(repo),
		repo:     repo,
	}
}

func (u *GenreUC) List(input entity.GenreFilters) (error, entity.GenreList) {
	return u.listUC.Execute(input)
}

func (u *GenreUC) Get(id int) (error, entity.Genre) {
	return u.getUC.Execute(id)
}

func (u *GenreUC) Create(input entity.Genre) (error, entity.Genre) {
	return u.createUC.Execute(input)
}

func (u *GenreUC) Update(id int, input entity.Genre) (error, entity.Genre) {
	return u.updateUC.Execute(id, input)
}

func (u *GenreUC) Delete(id int) error {
	return u.deleteUC.Execute(id)
}
