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

	var res []*models.Stores

	store, err := u.repo.GetStores(ctx)
	if err != nil {
		return nil, err
	}

	for _, data := range store {
		res = append(res, &models.Stores{
			StoreName:  data.StoreName,
			StorePhoto: data.StorePhoto,
		})
	}

	return res, nil
}
