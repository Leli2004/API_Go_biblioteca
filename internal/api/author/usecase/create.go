package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type CreateUC struct {
	repo author.Repository
}

func NewCreateUC(repo author.Repository) CreateUC {
	return CreateUC{repo: repo}
}

func (u *CreateUC) Execute(input entity.Author) (error, entity.Author) {
	err := input.Validate()
	if err != nil {
		return err, entity.Author{}
	}
	return u.repo.Create(input)
}
