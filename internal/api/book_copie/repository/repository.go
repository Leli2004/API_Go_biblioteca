package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type BookCopieRepo struct {
	list   ListRepo
	get    GetRepo
	create CreateRepo
	update UpdateRepo
	delete DeleteRepo
}

func NewRepository(db *sqlx.DB) *BookCopieRepo {
	return &BookCopieRepo{
		list:   NewListRepo(db),
		get:    NewGetRepo(db),
		create: NewCreateRepo(db),
		update: NewUpdateRepo(db),
		delete: NewDeleteRepo(db),
	}
}

func (r *BookCopieRepo) List(input entity.BookCopyFilters) (error, entity.BookCopyList) {
	return r.list.Execute(input)
}

func (r *BookCopieRepo) Get(id int) (error, entity.BookCopy) {
	return r.get.Execute(id)
}

func (r *BookCopieRepo) Create(input entity.BookCopy) (error, entity.BookCopy) {
	return r.create.Execute(input)
}

func (r *BookCopieRepo) Update(id int, input entity.BookCopy) (error, entity.BookCopy) {
	return r.update.Execute(id, input)
}

func (r *BookCopieRepo) Delete(id int) error {
	err, _ := r.delete.Execute(id)
	return err
}
