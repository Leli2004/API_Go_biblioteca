package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UpdateUC struct {
	repo publisher.Repository
}

func NewUpdateUC(repo publisher.Repository) UpdateUC {
	return UpdateUC{repo: repo}
}

func (u *UpdateUC) Execute(id int, input entity.Publisher) (error, entity.Publisher) {
	err := input.Validate()
	if err != nil {
		return err, entity.Publisher{}
	}
	return u.repo.Update(id, input)
}
