package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UpdateUC struct {
	repo genre.Repository
}

func NewUpdateUC(repo genre.Repository) UpdateUC {
	return UpdateUC{repo: repo}
}

func (u *UpdateUC) Execute(id int, input entity.Genre) (error, entity.Genre) {
	err := input.Validate()
	if err != nil {
		return err, entity.Genre{}
	}
	return u.repo.Update(id, input)
}
