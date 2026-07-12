package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UpdateUC struct {
	repo book.Repository
}

func NewUpdateUC(repo book.Repository) UpdateUC {
	return UpdateUC{repo: repo}
}

func (u *UpdateUC) Execute(id int, input entity.Book) (error, entity.Book) {
	err := input.Validate()
	if err != nil {
		return err, entity.Book{}
	}
	return u.repo.Update(id, input)
}
