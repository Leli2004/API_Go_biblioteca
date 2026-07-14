package repository

import (
	"context"
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

func NewRepository() *BookCopieRepo {
	return &BookCopieRepo{
		list:   NewListRepo(),
		get:    NewGetRepo(),
		create: NewCreateRepo(),
		update: NewUpdateRepo(),
		delete: NewDeleteRepo(),
	}
}

func (r *BookCopieRepo) List(ctx context.Context, tx *sqlx.Tx, input entity.BookCopyFilters) (context.Context, error, entity.BookCopyList) {
	return r.list.Execute(ctx, tx, input)
}

func (r *BookCopieRepo) Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.BookCopy) {
	return r.get.Execute(ctx, tx, id)
}

func (r *BookCopieRepo) Create(ctx context.Context, tx *sqlx.Tx, input entity.BookCopy) (context.Context, error, entity.BookCopy) {
	return r.create.Execute(ctx, tx, input)
}

func (r *BookCopieRepo) Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.BookCopy) (context.Context, error, entity.BookCopy) {
	return r.update.Execute(ctx, tx, id, input)
}

func (r *BookCopieRepo) Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error) {
	ctx, err, _ := r.delete.Execute(ctx, tx, id)
	return ctx, err
}
