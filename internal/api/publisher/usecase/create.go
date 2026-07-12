package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type CreateUC struct {
	repo publisher.Repository
}

func NewCreateUC(repo publisher.Repository) CreateUC {
	return CreateUC{repo: repo}
}

func (u *CreateUC) Execute(input entity.Publisher) (error, entity.Publisher) {
	err := input.Validate()
	if err != nil {
		return err, entity.Publisher{}
	}
	return u.repo.Create(input)
}
