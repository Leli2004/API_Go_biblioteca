package usecase

import (
	"context"
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type DeleteUC struct {
	db       *sqlx.DB
	repo     book.Repository
	redisCli *redis.Client
}

func NewDeleteUC(db *sqlx.DB, repo book.Repository, redisCli *redis.Client) DeleteUC {
	return DeleteUC{db: db, repo: repo, redisCli: redisCli}
}

func (u *DeleteUC) Execute(ctx context.Context, id int, claims *entity.AuthClaims) (returnedCtx context.Context, err error) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err
	}
	defer helpers.CloseTransaction(tx, &err)

	returnedCtx, err = u.repo.Delete(ctx, tx, id)
	if err != nil {
		return ctx, err
	}

	_ = u.redisCli.Del(ctx, fmt.Sprintf("biblioteca_book_get_%d", id), keyList).Err()

	return
}
