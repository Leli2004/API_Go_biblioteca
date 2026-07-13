package usecase

import (
	book_copie "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type BookCopieUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     book_copie.Repository
}

func NewUseCase(repo book_copie.Repository) BookCopieUC {
	return BookCopieUC{
		listUC:   NewListUC(repo),
		getUC:    NewGetUC(repo),
		createUC: NewCreateUC(repo),
		updateUC: NewUpdateUC(repo),
		deleteUC: NewDeleteUC(repo),
		repo:     repo,
	}
}

func (u *BookCopieUC) List(input entity.BookCopyFilters) (error, entity.BookCopyList) {
	return u.listUC.Execute(input)
}

func (u *BookCopieUC) Get(id int) (error, entity.BookCopy) {
	return u.getUC.Execute(id)
}

func (u *BookCopieUC) Create(input entity.BookCopy) (error, entity.BookCopy) {
	return u.createUC.Execute(input)
}

func (u *BookCopieUC) Update(id int, input entity.BookCopy) (error, entity.BookCopy) {
	return u.updateUC.Execute(id, input)
}

func (u *BookCopieUC) Delete(id int) error {
	return u.deleteUC.Execute(id)
}
