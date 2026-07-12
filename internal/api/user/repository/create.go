package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct {
	db *sqlx.DB
}

func NewCreateRepo(db *sqlx.DB) CreateRepo {
	return CreateRepo{db: db}
}

func (r *CreateRepo) Execute(input entity.User) (error, entity.User) {
	var user entity.User
	err := r.db.Get(&user, createSql, input.Name, input.Email, input.PasswordHash, input.Phone, input.Role, input.Active)
	if err != nil {
		return err, entity.User{}
	}

	return nil, user
}

var createSql = `
	INSERT INTO biblioteca.users (name, email, password_hash, phone, role, active)
	VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'member'), COALESCE($6, TRUE))
	RETURNING id, name, email, password_hash, phone, role, active, created_at, updated_at
`
