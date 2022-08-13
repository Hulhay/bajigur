package handlers

import (
	"context"
	"errors"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/shared"
)

func (h *handler) CreateStore(ctx context.Context, identifier string, params *models.StoresRequest) error {

	var (
		err  error
		role string
	)

	role, err = shared.GetRole(identifier)
	if err != nil {
		return err
	}
	if !shared.IsAllowedStoreAccountRole[role] {
		return errors.New("this account not allowed to access this endpoint")
	}

	err = h.useCase.CreateStore(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) GetStores(ctx context.Context, identifier string) ([]*models.Stores, error) {

	var (
		err  error
		role string
		res  []*models.Stores
	)

	role, err = shared.GetRole(identifier)
	if err != nil {
		return nil, err
	}
	if !shared.IsAllowedAccountRole[role] {
		return nil, errors.New("this account not allowed to access this endpoint")
	}

	res, err = h.useCase.GetStores(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *handler) GetStoreByID(ctx context.Context, identifier string, storeID string) (*models.Stores, error) {

	var (
		err  error
		role string
		res  *models.Stores
	)

	role, err = shared.GetRole(identifier)
	if err != nil {
		return nil, err
	}
	if !shared.IsAllowedAccountRole[role] {
		return nil, errors.New("this account not allowed to access this endpoint")
	}

	res, err = h.useCase.GetStoreByID(ctx, storeID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *handler) DeleteStoreByID(ctx context.Context, identifier string, storeID string) error {

	var (
		err  error
		role string
	)

	role, err = shared.GetRole(identifier)
	if err != nil {
		return err
	}
	if !shared.IsAllowedStoreAccountRole[role] {
		return errors.New("this account not allowed to access this endpoint")
	}

	err = h.useCase.DeleteStoreByID(ctx, storeID)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) UpdateStoreByID(ctx context.Context, identifier string, params *models.StoresRequest, storeID string) error {

	var (
		err  error
		role string
	)

	role, err = shared.GetRole(identifier)
	if err != nil {
		return err
	}
	if !shared.IsAllowedStoreAccountRole[role] {
		return errors.New("this account not allowed to access this endpoint")
	}

	err = h.useCase.UpdateStoreByID(ctx, params, storeID)
	if err != nil {
		return err
	}
	return nil
}
