package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type ListUC struct {
	repo book.Repository
}

func NewListUC(repo book.Repository) ListUC {
	return ListUC{repo: repo}
}

func (u *ListUC) Execute(input entity.BookFilters) (error, entity.BookList) {
	input.SetDefault()
	return u.repo.List(input)
}
