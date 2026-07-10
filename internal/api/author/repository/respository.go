package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthorRepo struct {
	list ListRepo
}

func NewRepository(db *sqlx.DB) *AuthorRepo {
	return &AuthorRepo{
		list: NewListRepo(db),
	}
}

func (r *AuthorRepo) List(input entity.AuthorFilters) (error, entity.AuthorList) {
	return r.list.Execute(input)
}
