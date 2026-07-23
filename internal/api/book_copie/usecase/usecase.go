package usecase

import (
	"context"

	book_copie "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type BookCopieUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     book_copie.Repository
}

func NewUseCase(db *sqlx.DB, repo book_copie.Repository) *BookCopieUC {
	return &BookCopieUC{
		listUC:   NewListUC(db, repo),
		getUC:    NewGetUC(db, repo),
		createUC: NewCreateUC(db, repo),
		updateUC: NewUpdateUC(db, repo),
		deleteUC: NewDeleteUC(db, repo),
		repo:     repo,
	}
}

func (u *BookCopieUC) List(ctx context.Context, input entity.BookCopyFilters) (context.Context, error, entity.BookCopyList) {
	return u.listUC.Execute(ctx, input)
}

func (u *BookCopieUC) Get(ctx context.Context, id int) (context.Context, error, entity.BookCopy) {
	return u.getUC.Execute(ctx, id)
}

func (u *BookCopieUC) Create(ctx context.Context, input entity.BookCopy, claims *entity.AuthClaims) (context.Context, error, entity.BookCopy) {
	return u.createUC.Execute(ctx, input, claims)
}

func (u *BookCopieUC) Update(ctx context.Context, id int, input entity.BookCopy, claims *entity.AuthClaims) (context.Context, error, entity.BookCopy) {
	return u.updateUC.Execute(ctx, id, input, claims)
}

func (u *BookCopieUC) Delete(ctx context.Context, id int, claims *entity.AuthClaims) (context.Context, error) {
	return u.deleteUC.Execute(ctx, id, claims)
}
