package usecase

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type UpdateUC struct {
	db       *sqlx.DB
	repo     author.Repository
	redisCli *redis.Client
}

func NewUpdateUC(db *sqlx.DB, repo author.Repository, redisCli *redis.Client) UpdateUC {
	return UpdateUC{db: db, repo: repo, redisCli: redisCli}
}

func (u *UpdateUC) Execute(ctx context.Context, id int, input entity.Author, claims *entity.AuthClaims) (returnedCtx context.Context, err error, result entity.Author) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	err = input.Validate()
	if err != nil {
		return ctx, err, entity.Author{}
	}

	returnedCtx, err, result = u.repo.Update(ctx, tx, id, input)
	if err != nil {
		return ctx, err, result
	}

	u.saveRedis(ctx, result, helpers.CacheGetKeyId(id, "author"))
	_ = u.redisCli.Del(ctx, keyList).Err()

	return
}

func (u *UpdateUC) saveRedis(ctx context.Context, result entity.Author, key string) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("error marshaling author to redis: %v", err)
	}

	err = u.redisCli.Set(ctx, key, data, 10*time.Minute).Err()
	if err != nil {
		log.Printf("error saving author to redis: %v", err)
	}
}
