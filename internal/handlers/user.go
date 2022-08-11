package handlers

import (
	"context"
	"hulhay-mall/internal/models"
)

func (h *handler) Register(ctx context.Context, params *models.RegisterRequest) error {
	err := h.useCase.Register(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) Login(ctx context.Context, params *models.LoginRequest) error {
	err := h.useCase.Login(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
