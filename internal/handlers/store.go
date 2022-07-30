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
