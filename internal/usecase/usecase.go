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
	GetStoreByID(ctx context.Context, storeID string) (*models.Stores, error)
	DeleteStoreByID(ctx context.Context, storeID string) error
	UpdateStoreByID(ctx context.Context, params *models.StoresRequest, stringID string) error
}

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
