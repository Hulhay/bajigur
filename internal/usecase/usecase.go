package usecase

import (
	"context"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/repositories"
)

type useCase struct {
	repo repositories.Repositories
}

type UseCase interface {
	CreateStore(ctx context.Context, params *models.StoresRequest) error
	GetStores(ctx context.Context) ([]*models.Stores, error)
}

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
