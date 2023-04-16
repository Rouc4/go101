package repository

import (
	"carshop/internal/models"
	"context"
)

type Store interface {
	Car() CarRepo
}

type CarRepo interface {
	Create(ctx context.Context, cc *models.CarObject) error   // для менеджеров
	ReadAll(ctx context.Context) ([]*models.CarObject, error) //
	Update(ctx context.Context, id uint)
	Delete(ctx context.Context, id uint)
}

/////////////////////////////////

/////////////////////////////////

/////////////////////////////////

/////////////////////////////////
