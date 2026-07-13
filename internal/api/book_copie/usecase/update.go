package usecase

import (
	book_copie "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UpdateUC struct {
	repo book_copie.Repository
}

func NewUpdateUC(repo book_copie.Repository) UpdateUC {
	return UpdateUC{repo: repo}
}

func (u *UpdateUC) Execute(id int, input entity.BookCopy) (error, entity.BookCopy) {
	err := input.Validate()
	if err != nil {
		return err, entity.BookCopy{}
	}
	return u.repo.Update(id, input)
}
