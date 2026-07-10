package usecase

import "github.com/Leli2004/API_Go_biblioteca/internal/api/genre"

type DeleteUC struct {
	repo genre.Repository
}

func NewDeleteUC(repo genre.Repository) DeleteUC {
	return DeleteUC{repo: repo}
}

func (u *DeleteUC) Execute(id int) error {
	return u.repo.Delete(id)
}
