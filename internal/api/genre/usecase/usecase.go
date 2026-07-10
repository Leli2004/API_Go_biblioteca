package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type GenreUC struct {
	listUC ListUC
	repo   genre.Repository
}

func NewUseCase(repo genre.Repository) GenreUC {
	return GenreUC{
		listUC: NewListUC(repo),
	}
}

func (u *GenreUC) List(input entity.GenreFilters) (error, entity.GenreList) {
	return u.listUC.Execute(input)
}
