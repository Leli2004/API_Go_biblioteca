package usecase

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     user.Repository
}

func NewUseCase(db *sqlx.DB, repo user.Repository) UserUC {
	return UserUC{
		listUC:   NewListUC(db, repo),
		getUC:    NewGetUC(db, repo),
		createUC: NewCreateUC(db, repo),
		updateUC: NewUpdateUC(db, repo),
		deleteUC: NewDeleteUC(db, repo),
		repo:     repo,
	}
}

func (u *UserUC) List(ctx context.Context, input entity.UserFilters) (context.Context, error, entity.UserList) {
	return u.listUC.Execute(ctx, input)
}

func (u *UserUC) Get(ctx context.Context, id int) (context.Context, error, entity.User) {
	return u.getUC.Execute(ctx, id)
}

func (u *UserUC) Create(ctx context.Context, input entity.User) (context.Context, error, entity.User) {
	return u.createUC.Execute(ctx, input)
}

func (u *UserUC) Update(ctx context.Context, id int, input entity.User) (context.Context, error, entity.User) {
	return u.updateUC.Execute(ctx, id, input)
}

func (u *UserUC) Delete(ctx context.Context, id int) (context.Context, error) {
	return u.deleteUC.Execute(ctx, id)
}
