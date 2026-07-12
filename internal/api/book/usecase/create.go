package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type CreateUC struct {
	repo book.Repository
}

func NewCreateUC(repo book.Repository) CreateUC {
	return CreateUC{repo: repo}
}

func (u *CreateUC) Execute(input entity.Book) (error, entity.Book) {
	err := input.Validate()
	if err != nil {
		return err, entity.Book{}
	}
	return u.repo.Create(input)
}
