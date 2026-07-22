package repository

import (
	"context"
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

func NewRepository() *AuthorRepo {
	return &AuthorRepo{
		list:   NewListRepo(),
		get:    NewGetRepo(),
		create: NewCreateRepo(),
		update: NewUpdateRepo(),
		delete: NewDeleteRepo(),
	}
}

func (r *AuthorRepo) List(ctx context.Context, tx *sqlx.Tx, input entity.AuthorFilters) (context.Context, error, entity.AuthorList) {
	return r.list.Execute(ctx, tx, input)
}

func (r *AuthorRepo) Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Author) {
	return r.get.Execute(ctx, tx, id)
}

func (r *AuthorRepo) Create(ctx context.Context, tx *sqlx.Tx, input entity.Author) (context.Context, error, entity.Author) {
	return r.create.Execute(ctx, tx, input)
}

func (r *AuthorRepo) Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.Author) (context.Context, error, entity.Author) {
	return r.update.Execute(ctx, tx, id, input)
}

func (r *AuthorRepo) Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error) {
	ctx, err, _ := r.delete.Execute(ctx, tx, id)
	return ctx, err
}
