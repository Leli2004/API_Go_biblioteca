package repository

import (
	"context"
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

func NewRepository() *PublisherRepo {
	return &PublisherRepo{
		list:   NewListRepo(),
		get:    NewGetRepo(),
		create: NewCreateRepo(),
		update: NewUpdateRepo(),
		delete: NewDeleteRepo(),
	}
}

func (r *PublisherRepo) List(ctx context.Context, tx *sqlx.Tx, input entity.PublisherFilters) (context.Context, error, entity.PublisherList) {
	return r.list.Execute(ctx, tx, input)
}

func (r *PublisherRepo) Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Publisher) {
	return r.get.Execute(ctx, tx, id)
}

func (r *PublisherRepo) Create(ctx context.Context, tx *sqlx.Tx, input entity.Publisher) (context.Context, error, entity.Publisher) {
	return r.create.Execute(ctx, tx, input)
}

func (r *PublisherRepo) Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Publisher) (context.Context, error, entity.Publisher) {
	return r.update.Execute(ctx, tx, id, input)
}

func (r *PublisherRepo) Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error) {
	ctx, err, _ := r.delete.Execute(ctx, tx, id)
	return ctx, err
}
