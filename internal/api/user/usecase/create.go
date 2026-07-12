package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type CreateUC struct {
	repo user.Repository
}

func NewCreateUC(repo user.Repository) CreateUC {
	return CreateUC{repo: repo}
}

func (u *CreateUC) Execute(input entity.User) (error, entity.User) {
	err := input.Validate()
	if err != nil {
		return err, entity.User{}
	}
	return u.repo.Create(input)
}
