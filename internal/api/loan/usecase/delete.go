package usecase

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type DeleteUC struct {
	db       *sqlx.DB
	repo     loan.Repository
	redisCli *redis.Client
}

func NewDeleteUC(db *sqlx.DB, repo loan.Repository, redisCli *redis.Client) DeleteUC {
	return DeleteUC{db: db, repo: repo, redisCli: redisCli}
}

func (u *DeleteUC) Execute(ctx context.Context, id int, claims *entity.AuthClaims) (returnedCtx context.Context, err error, result entity.Loan) {
	if err := security.ValidateRoles(claims, entity.RoleAdmin); err != nil {
		return ctx, err, result
	}

	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	returnedCtx, err, result = u.repo.Delete(ctx, tx, id)
	if err != nil {
		return ctx, err, result
	}

	_ = u.redisCli.Del(ctx, helpers.CacheGetKeyId(id, "loan"), keyList).Err()

	return
}
