package usecase

import (
	"context"
	"hulhay-mall/internal/models"
)

func (u *useCase) CreateStore(ctx context.Context, params *models.StoresRequest) error {

	err := u.repo.CreateStore(ctx, models.StoresRequest{
		StoreName:    params.StoreName,
		StorePhoto:   params.StorePhoto,
		Owner:        params.Owner,
		StorePhone:   params.StorePhone,
		StoreAddress: params.StoreAddress,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *useCase) GetStores(ctx context.Context) ([]*models.Stores, error) {

	res, err := u.repo.GetStores(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
