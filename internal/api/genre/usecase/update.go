package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type UpdateUC struct {
	db       *sqlx.DB
	repo     genre.Repository
	redisCli *redis.Client
}

func NewUpdateUC(db *sqlx.DB, repo genre.Repository, redisCli *redis.Client) UpdateUC {
	return UpdateUC{db: db, repo: repo, redisCli: redisCli}
}

func (u *UpdateUC) Execute(ctx context.Context, id int, input entity.Genre, claims *entity.AuthClaims) (returnedCtx context.Context, err error, result entity.Genre) {
	if err := security.ValidateRoles(claims, entity.RoleAdmin); err != nil {
		return ctx, err, result
	}

	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	err = input.Validate()
	if err != nil {
		return ctx, err, entity.Genre{}
	}

	returnedCtx, err, result = u.repo.Update(ctx, tx, id, input)
	if err != nil {
		return ctx, err, result
	}

	u.saveRedis(ctx, result, fmt.Sprintf("biblioteca_genre_get_%d", id))
	_ = u.redisCli.Del(ctx, keyList).Err()

	return
}

func (u *UpdateUC) saveRedis(ctx context.Context, result entity.Genre, key string) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("error marshaling genre to redis: %v", err)
	}

	err = u.redisCli.Set(ctx, key, data, 10*time.Minute).Err()
	if err != nil {
		log.Printf("error saving genre to redis: %v", err)
	}
}
