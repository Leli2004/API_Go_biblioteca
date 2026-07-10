package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type ListUC struct {
	repo genre.Repository
}

func NewListUC(repo genre.Repository) ListUC {
	return ListUC{
		repo: repo,
	}
}

func (u *ListUC) Execute(input entity.GenreFilters) (error, entity.GenreList) {
	input.SetDefault()
	return u.repo.List(input)
}
