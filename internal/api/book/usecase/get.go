package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type GetUC struct {
	db       *sqlx.DB
	repo     book.Repository
	redisCli *redis.Client
}

func NewGetUC(db *sqlx.DB, repo book.Repository, redisCli *redis.Client) GetUC {
	return GetUC{db: db, repo: repo, redisCli: redisCli}
}

func (u *GetUC) Execute(ctx context.Context, id int) (returnedCtx context.Context, err error, result entity.Book) {
	key := fmt.Sprintf("biblioteca_book_get_%d", id)

	cached, err := u.redisCli.Get(ctx, key).Result()
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

	returnedCtx, err, result = u.repo.Get(ctx, tx, id)
	if err != nil {
		return ctx, err, result
	}

	u.saveRedis(ctx, result, key)

	return
}

func (u *GetUC) saveRedis(ctx context.Context, result entity.Book, key string) {
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Errorf("Error saveRedis: %w", err)
	}

	err = u.redisCli.Set(ctx, key, data, 10*time.Minute).Err()
	if err != nil {
		log.Printf("error saving author to redis: %v", err)
	}
}
