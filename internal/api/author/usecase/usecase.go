package usecase

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthorUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     author.Repository
}

func NewUseCase(db *sqlx.DB, repo author.Repository) *AuthorUC {
	return &AuthorUC{
		listUC:   NewListUC(db, repo),
		getUC:    NewGetUC(db, repo),
		createUC: NewCreateUC(db, repo),
		updateUC: NewUpdateUC(db, repo),
		deleteUC: NewDeleteUC(db, repo),
		repo:     repo,
	}
}

func (u *AuthorUC) List(ctx context.Context, input entity.AuthorFilters) (context.Context, error, entity.AuthorList) {
	return u.listUC.Execute(ctx, input)
}

func (u *AuthorUC) Get(ctx context.Context, id int) (context.Context, error, entity.Author) {
	return u.getUC.Execute(ctx, id)
}

func (u *AuthorUC) Create(ctx context.Context, input entity.Author) (context.Context, error, entity.Author) {
	return u.createUC.Execute(ctx, input)
}

func (u *AuthorUC) Update(ctx context.Context, id int, input entity.Author) (context.Context, error, entity.Author) {
	return u.updateUC.Execute(ctx, id, input)
}

func (u *AuthorUC) Delete(ctx context.Context, id int) (context.Context, error) {
	return u.deleteUC.Execute(ctx, id)
}
