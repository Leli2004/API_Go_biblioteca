package usecase

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type ListUC struct {
	db       *sqlx.DB
	repo     book.Repository
	redisCli *redis.Client
}

func NewListUC(db *sqlx.DB, repo book.Repository, redisCli *redis.Client) ListUC {
	return ListUC{db: db, repo: repo, redisCli: redisCli}
}

const keyList = "biblioteca_book_list"

func (u *ListUC) Execute(ctx context.Context, input entity.BookFilters) (returnedCtx context.Context, err error, result entity.BookList) {

	cached, err := u.redisCli.Get(ctx, keyList).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cached), &result)
		if err == nil {
			return ctx, nil, result
		}
	}

	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	input.SetDefault()

	returnedCtx, err, result = u.repo.List(ctx, tx, input)
	if err != nil {
		return ctx, err, result
	}

	u.saveRedis(ctx, result, keyList)
	return
}

func (u *ListUC) saveRedis(ctx context.Context, result entity.BookList, key string) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("error marshaling book to redis: %v", err)
	}

	err = u.redisCli.Set(ctx, key, data, 10*time.Minute).Err()
	if err != nil {
		log.Printf("error saving book to redis: %v", err)
	}
}
