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
