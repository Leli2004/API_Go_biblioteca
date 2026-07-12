package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	list   ListRepo
	get    GetRepo
	create CreateRepo
	update UpdateRepo
	delete DeleteRepo
}

func NewRepository(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		list:   NewListRepo(db),
		get:    NewGetRepo(db),
		create: NewCreateRepo(db),
		update: NewUpdateRepo(db),
		delete: NewDeleteRepo(db),
	}
}

func (r *UserRepo) List(input entity.UserFilters) (error, entity.UserList) {
	return r.list.Execute(input)
}

func (r *UserRepo) Get(id int) (error, entity.User) {
	return r.get.Execute(id)
}

func (r *UserRepo) Create(input entity.User) (error, entity.User) {
	return r.create.Execute(input)
}

func (r *UserRepo) Update(id int, input entity.User) (error, entity.User) {
	return r.update.Execute(id, input)
}

func (r *UserRepo) Delete(id int) error {
	err, _ := r.delete.Execute(id)
	return err
}
