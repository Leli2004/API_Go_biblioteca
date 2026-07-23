package repository

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	list           ListRepo
	get            GetRepo
	create         CreateRepo
	update         UpdateRepo
	delete         DeleteRepo
	getByUsername  GetByUsernameRepo
	usernameExists UsernameExistsRepo
	emailExists    EmailExistsRepo
}

func NewRepository() *UserRepo {
	return &UserRepo{
		list:           NewListRepo(),
		get:            NewGetRepo(),
		create:         NewCreateRepo(),
		update:         NewUpdateRepo(),
		delete:         NewDeleteRepo(),
		getByUsername:  NewGetByUsernameRepo(),
		usernameExists: NewUsernameExistsRepo(),
		emailExists:    NewEmailExistsRepo(),
	}
}

func (r *UserRepo) List(ctx context.Context, tx *sqlx.Tx, input entity.UserFilters) (context.Context, error, entity.UserList) {
	return r.list.Execute(ctx, tx, input)
}

func (r *UserRepo) Get(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.User) {
	return r.get.Execute(ctx, tx, id)
}

func (r *UserRepo) Create(ctx context.Context, tx *sqlx.Tx, input entity.User) (context.Context, error, entity.User) {
	return r.create.Execute(ctx, tx, input)
}

func (r *UserRepo) Update(ctx context.Context, tx *sqlx.Tx, id int, input entity.User) (context.Context, error, entity.User) {
	return r.update.Execute(ctx, tx, id, input)
}

func (r *UserRepo) Delete(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error) {
	ctx, err, _ := r.delete.Execute(ctx, tx, id)
	return ctx, err
}

func (r *UserRepo) GetByUsername(ctx context.Context, tx *sqlx.Tx, username string) (context.Context, error, entity.User) {
	return r.getByUsername.Execute(ctx, tx, username)
}

func (r *UserRepo) UsernameExists(ctx context.Context, tx *sqlx.Tx, username string, excludeUserID int) (context.Context, error, bool) {
	return r.usernameExists.Execute(ctx, tx, username, excludeUserID)
}

func (r *UserRepo) EmailExists(ctx context.Context, tx *sqlx.Tx, email string, excludeUserID int) (context.Context, error, bool) {
	return r.emailExists.Execute(ctx, tx, email, excludeUserID)
}
