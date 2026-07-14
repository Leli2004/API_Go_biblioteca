package repository

import (
	"context"
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

func NewRepository() *GenreRepo {
	return &GenreRepo{
		list:   NewListRepo(),
		get:    NewGetRepo(),
		create: NewCreateRepo(),
		update: NewUpdateRepo(),
		delete: NewDeleteRepo(),
	}
}

func (r *GenreRepo) List(ctx context.Context, tx *sqlx.Tx, input entity.GenreFilters) (context.Context, error, entity.GenreList) {
	return r.list.Execute(ctx, tx, input)
}

func (r *GenreRepo) Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Genre) {
	return r.get.Execute(ctx, tx, id)
}

func (r *GenreRepo) Create(ctx context.Context, tx *sqlx.Tx, input entity.Genre) (context.Context, error, entity.Genre) {
	return r.create.Execute(ctx, tx, input)
}

func (r *GenreRepo) Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Genre) (context.Context, error, entity.Genre) {
	return r.update.Execute(ctx, tx, id, input)
}

func (r *GenreRepo) Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error) {
	ctx, err, _ := r.delete.Execute(ctx, tx, id)
	return ctx, err
}
