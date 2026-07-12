package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UpdateUC struct {
	repo user.Repository
}

func NewUpdateUC(repo user.Repository) UpdateUC {
	return UpdateUC{repo: repo}
}

func (u *UpdateUC) Execute(id int, input entity.User) (error, entity.User) {
	err := input.Validate()
	if err != nil {
		return err, entity.User{}
	}
	return u.repo.Update(id, input)
}
