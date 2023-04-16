package service

import (
	"carshop/internal/models"
	"context"
)

type CarService interface {
	GetCar(ctx context.Context) (*models.CarObject, error)
	CreateCar(ctx context.Context, c *models.CarObject) error
}
