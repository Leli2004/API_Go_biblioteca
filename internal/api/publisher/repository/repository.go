package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PublisherRepo struct {
	list   ListRepo
	get    GetRepo
	create CreateRepo
	update UpdateRepo
	delete DeleteRepo
}

func NewRepository(db *sqlx.DB) *PublisherRepo {
	return &PublisherRepo{
		list:   NewListRepo(db),
		get:    NewGetRepo(db),
		create: NewCreateRepo(db),
		update: NewUpdateRepo(db),
		delete: NewDeleteRepo(db),
	}
}

func (r *PublisherRepo) List(input entity.PublisherFilters) (error, entity.PublisherList) {
	return r.list.Execute(input)
}

func (r *PublisherRepo) Get(id int) (error, entity.Publisher) {
	return r.get.Execute(id)
}

func (r *PublisherRepo) Create(input entity.Publisher) (error, entity.Publisher) {
	return r.create.Execute(input)
}

func (r *PublisherRepo) Update(id int, input entity.Publisher) (error, entity.Publisher) {
	return r.update.Execute(id, input)
}

func (r *PublisherRepo) Delete(id int) error {
	err, _ := r.delete.Execute(id)
	return err
}
