package usecase

import (
	"context"
	"hulhay-mall/internal/models"
)

func (u *useCase) CreateStore(ctx context.Context, params *models.StoresRequest) error {

	errRepoStore := u.repo.CreateStore(ctx, models.StoresRequest{
		StoreName:    params.StoreName,
		StorePhoto:   params.StorePhoto,
		Owner:        params.Owner,
		StorePhone:   params.StorePhone,
		StoreAddress: params.StoreAddress,
	})

	if errRepoStore != nil {
		return errRepoStore
	}

	return nil
}
