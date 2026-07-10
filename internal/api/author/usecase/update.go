package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UpdateUC struct {
	repo author.Repository
}

func NewUpdateUC(repo author.Repository) UpdateUC {
	return UpdateUC{repo: repo}
}

func (u *UpdateUC) Execute(id int, input entity.Author) (error, entity.Author) {
	err := input.Validate()
	if err != nil {
		return err, entity.Author{}
	}
	return u.repo.Update(id, input)
}
