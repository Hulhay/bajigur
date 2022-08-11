package usecase

import (
	"context"
	"hulhay-mall/internal/models"
)

func (u *useCase) CreateStore(ctx context.Context, params *models.StoresRequest) error {

	err := u.repo.CreateStore(ctx, &models.StoresRequest{
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

func (u *useCase) GetStoreByID(ctx context.Context, storeID string) (*models.Stores, error) {

	store, err := u.repo.GetStoreByID(ctx, storeID)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (u *useCase) DeleteStoreByID(ctx context.Context, storeID string) error {

	_, err := u.repo.GetStoreByID(ctx, storeID)
	if err != nil {
		return err
	}

	err = u.repo.DeleteStoreByID(ctx, storeID)
	if err != nil {
		return err
	}
	return nil
}

func (u *useCase) UpdateStoreByID(ctx context.Context, params *models.StoresRequest, storeID string) error {

	_, err := u.repo.GetStoreByID(ctx, storeID)
	if err != nil {
		return err
	}

	req := &models.StoresRequest{}
	req.StoreName = params.StoreName
	req.StorePhoto = params.StorePhoto
	req.Owner = params.Owner
	req.StorePhone = params.StorePhone
	req.StoreAddress = params.StoreAddress

	err = u.repo.UpdateStoreByID(ctx, req, storeID)
	if err != nil {
		return err
	}
	return nil
}
