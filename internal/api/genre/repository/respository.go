package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GenreRepo struct {
	list ListRepo
}

func NewRepository(db *sqlx.DB) *GenreRepo {
	return &GenreRepo{
		list: NewListRepo(db),
	}
}

func (r *GenreRepo) List(input entity.GenreFilters) (error, entity.GenreList) {
	return r.list.Execute(input)
}
