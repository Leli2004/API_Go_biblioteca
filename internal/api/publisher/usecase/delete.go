package usecase

import "github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"

type DeleteUC struct {
	repo publisher.Repository
}

func NewDeleteUC(repo publisher.Repository) DeleteUC {
	return DeleteUC{repo: repo}
}

func (u *DeleteUC) Execute(id int) error {
	return u.repo.Delete(id)
}
