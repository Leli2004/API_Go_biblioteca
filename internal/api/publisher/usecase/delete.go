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

type DeleteUC struct {
	db       *sqlx.DB
	repo     publisher.Repository
	redisCli *redis.Client
}

func NewDeleteUC(db *sqlx.DB, repo publisher.Repository, redisCli *redis.Client) DeleteUC {
	return DeleteUC{db: db, repo: repo, redisCli: redisCli}
}

func (u *DeleteUC) Execute(ctx context.Context, id int, claims *entity.AuthClaims) (returnedCtx context.Context, err error) {
	if err := security.ValidateRoles(claims, entity.RoleAdmin); err != nil {
		return ctx, err
	}

	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err
	}
	defer helpers.CloseTransaction(tx, &err)

	returnedCtx, err = u.repo.Delete(ctx, tx, id)
	if err != nil {
		return ctx, err
	}

	_ = u.redisCli.Del(ctx, helpers.CacheGetKeyId(id, "publisher"), keyList).Err()

	return
}
