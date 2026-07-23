package usecase

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
)

type UpdateUC struct {
	db   *sqlx.DB
	repo publisher.Repository
}

func NewUpdateUC(db *sqlx.DB, repo publisher.Repository) UpdateUC {
	return UpdateUC{db: db, repo: repo}
}

func (u *UpdateUC) Execute(ctx context.Context, id int, input entity.Publisher, claims *entity.AuthClaims) (returnedCtx context.Context, err error, result entity.Publisher) {
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

	return u.repo.Update(ctx, tx, id, input)
}
