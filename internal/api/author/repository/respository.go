package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthorRepo struct {
	list   ListRepo
	get    GetRepo
	create CreateRepo
	update UpdateRepo
	delete DeleteRepo
}

func NewRepository(db *sqlx.DB) *AuthorRepo {
	return &AuthorRepo{
		list:   NewListRepo(db),
		get:    NewGetRepo(db),
		create: NewCreateRepo(db),
		update: NewUpdateRepo(db),
		delete: NewDeleteRepo(db),
	}
}

func (r *AuthorRepo) List(input entity.AuthorFilters) (error, entity.AuthorList) {
	return r.list.Execute(input)
}

func (r *AuthorRepo) Get(id int) (error, entity.Author) {
	return r.get.Execute(id)
}

func (r *AuthorRepo) Create(input entity.Author) (error, entity.Author) {
	return r.create.Execute(input)
}

func (r *AuthorRepo) Update(id int, input entity.Author) (error, entity.Author) {
	return r.update.Execute(id, input)
}

func (r *AuthorRepo) Delete(id int) error {
	err, _ := r.delete.Execute(id)
	return err
}
