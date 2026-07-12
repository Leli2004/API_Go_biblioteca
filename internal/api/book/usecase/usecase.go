package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type BookUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     book.Repository
}

func NewUseCase(repo book.Repository) BookUC {
	return BookUC{
		listUC:   NewListUC(repo),
		getUC:    NewGetUC(repo),
		createUC: NewCreateUC(repo),
		updateUC: NewUpdateUC(repo),
		deleteUC: NewDeleteUC(repo),
		repo:     repo,
	}
}

func (u *BookUC) List(input entity.BookFilters) (error, entity.BookList) {
	return u.listUC.Execute(input)
}

func (u *BookUC) Get(id int) (error, entity.Book) {
	return u.getUC.Execute(id)
}

func (u *BookUC) Create(input entity.Book) (error, entity.Book) {
	return u.createUC.Execute(input)
}

func (u *BookUC) Update(id int, input entity.Book) (error, entity.Book) {
	return u.updateUC.Execute(id, input)
}

func (u *BookUC) Delete(id int) error {
	return u.deleteUC.Execute(id)
}
