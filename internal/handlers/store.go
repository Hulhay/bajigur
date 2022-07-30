package handlers

import (
	"context"
	"hulhay-mall/internal/models"
)

func (h *handler) CreateStore(ctx context.Context, params *models.StoresRequest) error {
	err := h.useCase.CreateStore(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) GetStores(ctx context.Context) ([]*models.Stores, error) {
	res, err := h.useCase.GetStores(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *handler) GetStoreByID(ctx context.Context, storeID string) (*models.Stores, error) {
	res, err := h.useCase.GetStoreByID(ctx, storeID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *handler) DeleteStoreByID(ctx context.Context, storeID string) error {
	err := h.useCase.DeleteStoreByID(ctx, storeID)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) UpdateStoreByID(ctx context.Context, params *models.StoresRequest, storeID string) error {
	err := h.useCase.UpdateStoreByID(ctx, params, storeID)
	if err != nil {
		return err
	}
	return nil
}
