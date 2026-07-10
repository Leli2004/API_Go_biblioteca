package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type CreateUC struct {
	repo genre.Repository
}

func NewCreateUC(repo genre.Repository) CreateUC {
	return CreateUC{repo: repo}
}

func (u *CreateUC) Execute(input entity.Genre) (error, entity.Genre) {
	err := input.Validate()
	if err != nil {
		return err, entity.Genre{}
	}
	return u.repo.Create(input)
}
