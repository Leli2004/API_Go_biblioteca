package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UserUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     user.Repository
}

func NewUseCase(repo user.Repository) UserUC {
	return UserUC{
		listUC:   NewListUC(repo),
		getUC:    NewGetUC(repo),
		createUC: NewCreateUC(repo),
		updateUC: NewUpdateUC(repo),
		deleteUC: NewDeleteUC(repo),
		repo:     repo,
	}
}

func (u *UserUC) List(input entity.UserFilters) (error, entity.UserList) {
	return u.listUC.Execute(input)
}

func (u *UserUC) Get(id int) (error, entity.User) {
	return u.getUC.Execute(id)
}

func (u *UserUC) Create(input entity.User) (error, entity.User) {
	return u.createUC.Execute(input)
}

func (u *UserUC) Update(id int, input entity.User) (error, entity.User) {
	return u.updateUC.Execute(id, input)
}

func (u *UserUC) Delete(id int) error {
	return u.deleteUC.Execute(id)
}
