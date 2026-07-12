package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	list   ListRepo
	get    GetRepo
	create CreateRepo
	update UpdateRepo
	delete DeleteRepo
}

func NewRepository(db *sqlx.DB) *BookRepo {
	return &BookRepo{
		list:   NewListRepo(db),
		get:    NewGetRepo(db),
		create: NewCreateRepo(db),
		update: NewUpdateRepo(db),
		delete: NewDeleteRepo(db),
	}
}

func (r *BookRepo) List(input entity.BookFilters) (error, entity.BookList) {
	return r.list.Execute(input)
}

func (r *BookRepo) Get(id int) (error, entity.Book) {
	return r.get.Execute(id)
}

func (r *BookRepo) Create(input entity.Book) (error, entity.Book) {
	return r.create.Execute(input)
}

func (r *BookRepo) Update(id int, input entity.Book) (error, entity.Book) {
	return r.update.Execute(id, input)
}

func (r *BookRepo) Delete(id int) error {
	err, _ := r.delete.Execute(id)
	return err
}
