package memo

import (
	"carshop/internal/models"
	"carshop/internal/repository"
	"context"
)

type Repo struct {
	data map[uint8]models.CarObject
}

func New() repository.CarRepo {
	return &Repo{}
}

func (r *Repo) Create(ctx context.Context, cc *models.CarObject) error {
	r.data[cc.ID] = *cc
	return nil
}

func (r *Repo) ReadAll(ctx context.Context) ([]*models.CarObject, error) {
	var res []*models.CarObject

	for _, v := range r.data {
		res = append(res, &v)
	}

	return res, nil
}

func (r *Repo) Update(ctx context.Context, id uint) {
	panic("implement me")
}

func (r *Repo) Delete(ctx context.Context, id uint) {
	panic("implement me")
}
