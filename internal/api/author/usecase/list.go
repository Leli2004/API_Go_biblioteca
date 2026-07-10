package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type ListUC struct {
	repo author.Repository
}

func NewListUC(repo author.Repository) ListUC {
	return ListUC{
		repo: repo,
	}
}

func (u *ListUC) Execute(input entity.AuthorFilters) (error, entity.AuthorList) {
	input.SetDefault()
	return u.repo.List(input)
}
