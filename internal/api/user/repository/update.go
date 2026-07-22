package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct{}

func NewUpdateRepo() UpdateRepo {
	return UpdateRepo{}
}

func (r *UpdateRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int, input entity.User) (context.Context, error, entity.User) {
	var user entity.User
	err := tx.GetContext(ctx, &user, updateSql, input.Name, input.Email, input.Username, input.PasswordHash, input.Phone, input.Role, input.Active, id)
	if err != nil {
		return ctx, err, entity.User{}
	}

	return ctx, nil, user
}

var updateSql = `
	UPDATE biblioteca.users
	SET name = $1,
	    email = $2,
		username = $3,
	    password_hash = $4,
	    phone = $5,
	    role = COALESCE(NULLIF($6, ''), role),
	    active = COALESCE($7, active),
	    updated_at = NOW()
	WHERE id = $8
	RETURNING id, name, email, password_hash, phone, role, active, created_at, updated_at
`
