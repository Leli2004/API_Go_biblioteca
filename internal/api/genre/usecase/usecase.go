package usecase

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GenreUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     genre.Repository
}

func NewUseCase(db *sqlx.DB, repo genre.Repository) *GenreUC {
	return &GenreUC{
		listUC:   NewListUC(db, repo),
		getUC:    NewGetUC(db, repo),
		createUC: NewCreateUC(db, repo),
		updateUC: NewUpdateUC(db, repo),
		deleteUC: NewDeleteUC(db, repo),
		repo:     repo,
	}
}

func (u *GenreUC) List(ctx context.Context, input entity.GenreFilters) (context.Context, error, entity.GenreList) {
	return u.listUC.Execute(ctx, input)
}

func (u *GenreUC) Get(ctx context.Context, id int) (context.Context, error, entity.Genre) {
	return u.getUC.Execute(ctx, id)
}

func (u *GenreUC) Create(ctx context.Context, input entity.Genre, claims *entity.AuthClaims) (context.Context, error, entity.Genre) {
	return u.createUC.Execute(ctx, input, claims)
}

func (u *GenreUC) Update(ctx context.Context, id int, input entity.Genre, claims *entity.AuthClaims) (context.Context, error, entity.Genre) {
	return u.updateUC.Execute(ctx, id, input, claims)
}

func (u *GenreUC) Delete(ctx context.Context, id int, claims *entity.AuthClaims) (context.Context, error) {
	return u.deleteUC.Execute(ctx, id, claims)
}
