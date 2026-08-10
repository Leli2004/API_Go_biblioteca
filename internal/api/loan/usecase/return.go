package usecase

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type ReturnUC struct {
	db       *sqlx.DB
	repo     loan.Repository
	redisCli *redis.Client
}

func NewReturnUC(db *sqlx.DB, repo loan.Repository, redisCli *redis.Client) ReturnUC {
	return ReturnUC{db: db, repo: repo, redisCli: redisCli}
}

func (u *ReturnUC) Execute(ctx context.Context, loanId int, returnedAt *string, claims *entity.AuthClaims) (returnedCtx context.Context, err error, result entity.Loan) {
	if err := security.ValidateRoles(claims, entity.RoleAdmin); err != nil {
		return ctx, err, result
	}

	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	ctx, err, ln := u.repo.Get(ctx, tx, loanId)
	if err != nil {
		return ctx, err, entity.Loan{}
	}

	if returnedAt == nil || *returnedAt == "" {
		t := time.Now().UTC().Format(time.RFC3339)
		returnedAt = &t
	}

	ln.ReturnedAt = returnedAt
	ln.Status = "returned"

	returnedCtx, err, result = u.repo.Update(ctx, tx, loanId, ln)
	if err != nil {
		return ctx, err, result
	}

	u.saveRedis(ctx, result, helpers.CacheGetKeyId(loanId, "loan"))
	_ = u.redisCli.Del(ctx, keyList).Err()

	return
}

func (u *ReturnUC) saveRedis(ctx context.Context, result entity.Loan, key string) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("error marshaling author to redis: %v", err)
	}

	err = u.redisCli.Set(ctx, key, data, 10*time.Minute).Err()
	if err != nil {
		log.Printf("error saving author to redis: %v", err)
	}
}
