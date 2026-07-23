package usecase

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PublisherUC struct {
	listUC   ListUC
	getUC    GetUC
	createUC CreateUC
	updateUC UpdateUC
	deleteUC DeleteUC
	repo     publisher.Repository
}

func NewUseCase(db *sqlx.DB, repo publisher.Repository) *PublisherUC {
	return &PublisherUC{
		listUC:   NewListUC(db, repo),
		getUC:    NewGetUC(db, repo),
		createUC: NewCreateUC(db, repo),
		updateUC: NewUpdateUC(db, repo),
		deleteUC: NewDeleteUC(db, repo),
		repo:     repo,
	}
}

func (u *PublisherUC) List(ctx context.Context, input entity.PublisherFilters) (context.Context, error, entity.PublisherList) {
	return u.listUC.Execute(ctx, input)
}

func (u *PublisherUC) Get(ctx context.Context, id int) (context.Context, error, entity.Publisher) {
	return u.getUC.Execute(ctx, id)
}

func (u *PublisherUC) Create(ctx context.Context, input entity.Publisher, claims *entity.AuthClaims) (context.Context, error, entity.Publisher) {
	return u.createUC.Execute(ctx, input, claims)
}

func (u *PublisherUC) Update(ctx context.Context, id int, input entity.Publisher, claims *entity.AuthClaims) (context.Context, error, entity.Publisher) {
	return u.updateUC.Execute(ctx, id, input, claims)
}

func (u *PublisherUC) Delete(ctx context.Context, id int, claims *entity.AuthClaims) (context.Context, error) {
	return u.deleteUC.Execute(ctx, id, claims)
}
