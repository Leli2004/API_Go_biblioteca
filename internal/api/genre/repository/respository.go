package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GenreRepo struct {
	list   ListRepo
	get    GetRepo
	create CreateRepo
	update UpdateRepo
	delete DeleteRepo
}

func NewRepository(db *sqlx.DB) *GenreRepo {
	return &GenreRepo{
		list:   NewListRepo(db),
		get:    NewGetRepo(db),
		create: NewCreateRepo(db),
		update: NewUpdateRepo(db),
		delete: NewDeleteRepo(db),
	}
}

func (r *GenreRepo) List(input entity.GenreFilters) (error, entity.GenreList) {
	return r.list.Execute(input)
}

func (r *GenreRepo) Get(id int) (error, entity.Genre) {
	return r.get.Execute(id)
}

func (r *GenreRepo) Create(input entity.Genre) (error, entity.Genre) {
	return r.create.Execute(input)
}

func (r *GenreRepo) Update(id int, input entity.Genre) (error, entity.Genre) {
	return r.update.Execute(id, input)
}

func (r *GenreRepo) Delete(id int) error {
	err, _ := r.delete.Execute(id)
	return err
}
