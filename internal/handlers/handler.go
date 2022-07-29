package handlers

import (
	configs "hulhay-mall/config"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	GetHealtcheck() (*models.Health, error)
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
