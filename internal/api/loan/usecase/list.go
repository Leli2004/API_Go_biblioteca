package usecase

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

const keyList = "biblioteca_publisher_list"

type ListUC struct {
	db       *sqlx.DB
	repo     loan.Repository
	redisCli *redis.Client
}

func NewListUC(db *sqlx.DB, repo loan.Repository, redisCli *redis.Client) ListUC {
	return ListUC{db: db, repo: repo, redisCli: redisCli}
}

func (u *ListUC) Execute(ctx context.Context, input entity.LoanFilters) (returnedCtx context.Context, err error, result entity.LoanList) {
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

func (u *ListUC) saveRedis(ctx context.Context, result entity.LoanList, key string) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("error marshaling author to redis: %v", err)
	}

	err = u.redisCli.Set(ctx, key, data, 10*time.Minute).Err()
	if err != nil {
		log.Printf("error saving author to redis: %v", err)
	}
}
