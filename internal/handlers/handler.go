package handlers

import (
	"context"
	configs "hulhay-mall/config"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	GetHealtcheck() (*models.Health, error)

	CreateStore(ctx context.Context, params *models.StoresRequest) error
	GetStores(ctx context.Context) ([]*models.Stores, error)
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
