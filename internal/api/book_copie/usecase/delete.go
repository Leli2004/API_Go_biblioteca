package usecase

import (
	"context"

	book_copie "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
)

type DeleteUC struct {
	db   *sqlx.DB
	repo book_copie.Repository
}

func NewDeleteUC(db *sqlx.DB, repo book_copie.Repository) DeleteUC {
	return DeleteUC{db: db, repo: repo}
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

	return u.repo.Delete(ctx, tx, id)
}
