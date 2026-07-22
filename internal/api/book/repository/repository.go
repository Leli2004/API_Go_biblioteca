package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	list    ListRepo
	get     GetRepo
	sublist SublistRepo
	create  CreateRepo
	update  UpdateRepo
	delete  DeleteRepo
}

func NewRepository() *BookRepo {
	sublist := NewSublistRepo()
	return &BookRepo{
		list:    NewListRepo(),
		sublist: sublist,
		get:     NewGetRepo(sublist),
		create:  NewCreateRepo(),
		update:  NewUpdateRepo(),
		delete:  NewDeleteRepo(),
	}
}

func (r *BookRepo) List(ctx context.Context, tx *sqlx.Tx, input entity.BookFilters) (context.Context, error, entity.BookList) {
	return r.list.Execute(ctx, tx, input)
}

func (r *BookRepo) Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Book) {
	return r.get.Execute(ctx, tx, id)
}

func (r *BookRepo) Create(ctx context.Context, tx *sqlx.Tx, input entity.Book) (context.Context, error, entity.Book) {
	return r.create.Execute(ctx, tx, input)
}

func (r *BookRepo) Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Book) (context.Context, error, entity.Book) {
	return r.update.Execute(ctx, tx, id, input)
}

func (r *BookRepo) Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error) {
	ctx, err, _ := r.delete.Execute(ctx, tx, id)
	return ctx, err
}
