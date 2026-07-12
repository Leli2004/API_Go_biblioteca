package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type PublisherUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     publisher.Repository
}

func NewUseCase(repo publisher.Repository) PublisherUC {
	return PublisherUC{
		listUC:   NewListUC(repo),
		getUC:    NewGetUC(repo),
		createUC: NewCreateUC(repo),
		updateUC: NewUpdateUC(repo),
		deleteUC: NewDeleteUC(repo),
		repo:     repo,
	}
}

func (u *PublisherUC) List(input entity.PublisherFilters) (error, entity.PublisherList) {
	return u.listUC.Execute(input)
}

func (u *PublisherUC) Get(id int) (error, entity.Publisher) {
	return u.getUC.Execute(id)
}

func (u *PublisherUC) Create(input entity.Publisher) (error, entity.Publisher) {
	return u.createUC.Execute(input)
}

func (u *PublisherUC) Update(id int, input entity.Publisher) (error, entity.Publisher) {
	return u.updateUC.Execute(id, input)
}

func (u *PublisherUC) Delete(id int) error {
	return u.deleteUC.Execute(id)
}
