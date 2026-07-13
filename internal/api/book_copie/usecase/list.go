package usecase

import (
	book_copie "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type ListUC struct {
	repo book_copie.Repository
}

func NewListUC(repo book_copie.Repository) ListUC {
	return ListUC{repo: repo}
}

func (u *ListUC) Execute(input entity.BookCopyFilters) (error, entity.BookCopyList) {
	input.SetDefault()
	return u.repo.List(input)
}
