package usecase

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type CreateUC struct {
	db       *sqlx.DB
	repo     publisher.Repository
	redisCli *redis.Client
}

func NewCreateUC(db *sqlx.DB, repo publisher.Repository, redisCli *redis.Client) CreateUC {
	return CreateUC{db: db, repo: repo, redisCli: redisCli}
}

func (u *CreateUC) Execute(ctx context.Context, input entity.Publisher, claims *entity.AuthClaims) (returnedCtx context.Context, err error, result entity.Publisher) {
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
		return ctx, err, entity.Publisher{}
	}

	returnedCtx, err, result = u.repo.Create(ctx, tx, input)
	if err != nil {
		return ctx, err, result
	}

	_ = u.redisCli.Del(ctx, keyList).Err()

	return
}
