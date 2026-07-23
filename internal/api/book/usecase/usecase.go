package usecase

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type BookUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     book.Repository
	redisCli *redis.Client
}

func NewUseCase(db *sqlx.DB, repo book.Repository, redisCli *redis.Client) *BookUC {
	return &BookUC{
		listUC:   NewListUC(db, repo, redisCli),
		getUC:    NewGetUC(db, repo, redisCli),
		createUC: NewCreateUC(db, repo),
		updateUC: NewUpdateUC(db, repo),
		deleteUC: NewDeleteUC(db, repo),
		repo:     repo,
		redisCli: redisCli,
	}
}

func (u *BookUC) List(ctx context.Context, input entity.BookFilters) (context.Context, error, entity.BookList) {
	return u.listUC.Execute(ctx, input)
}

func (u *BookUC) Get(ctx context.Context, id int) (context.Context, error, entity.Book) {
	return u.getUC.Execute(ctx, id)
}

func (u *BookUC) Create(ctx context.Context, input entity.Book, claims *entity.AuthClaims) (context.Context, error, entity.Book) {
	return u.createUC.Execute(ctx, input, claims)
}

func (u *BookUC) Update(ctx context.Context, id int, input entity.Book, claims *entity.AuthClaims) (context.Context, error, entity.Book) {
	return u.updateUC.Execute(ctx, id, input, claims)
}

func (u *BookUC) Delete(ctx context.Context, id int, claims *entity.AuthClaims) (context.Context, error) {
	return u.deleteUC.Execute(ctx, id, claims)
}
