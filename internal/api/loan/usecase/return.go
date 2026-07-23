package usecase

import (
	"context"
	"time"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
)

type ReturnUC struct {
	db   *sqlx.DB
	repo loan.Repository
}

func NewReturnUC(db *sqlx.DB, repo loan.Repository) ReturnUC {
	return ReturnUC{db: db, repo: repo}
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

	ctx, err, updated := u.repo.Update(ctx, tx, loanId, ln)
	if err != nil {
		return ctx, err, entity.Loan{}
	}

	return ctx, nil, updated
}
