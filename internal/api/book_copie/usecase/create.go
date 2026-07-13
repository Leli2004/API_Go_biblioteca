package usecase

import (
	book_copie "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type CreateUC struct {
	repo book_copie.Repository
}

func NewCreateUC(repo book_copie.Repository) CreateUC {
	return CreateUC{repo: repo}
}

func (u *CreateUC) Execute(input entity.BookCopy) (error, entity.BookCopy) {
	err := input.Validate()
	if err != nil {
		return err, entity.BookCopy{}
	}
	return u.repo.Create(input)
}
