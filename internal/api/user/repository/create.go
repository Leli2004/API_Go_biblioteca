package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct{}

func NewCreateRepo() CreateRepo {
	return CreateRepo{}
}

func (r *CreateRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.User) (context.Context, error, entity.User) {
	var user entity.User
	err := tx.GetContext(ctx, &user, createSql, input.Name, input.Email, input.PasswordHash, input.Phone, input.Role, input.Active)
	if err != nil {
		return ctx, err, entity.User{}
	}

	return ctx, nil, user
}

var createSql = `
	INSERT INTO biblioteca.users (name, email, password_hash, phone, role, active)
	VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'member'), COALESCE($6, TRUE))
	RETURNING id, name, email, password_hash, phone, role, active, created_at, updated_at
`
