package usecase

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type ListUC struct {
	repo publisher.Repository
}

func NewListUC(repo publisher.Repository) ListUC {
	return ListUC{repo: repo}
}

func (u *ListUC) Execute(input entity.PublisherFilters) (error, entity.PublisherList) {
	input.SetDefault()
	return u.repo.List(input)
}
