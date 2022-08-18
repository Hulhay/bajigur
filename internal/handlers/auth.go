package handlers

import (
	"context"
	"hulhay-mall/internal/apis/operations/user"
	"hulhay-mall/internal/models"
)

func (h *handler) Register(ctx context.Context, params *models.RegisterRequest) error {
	err := h.useCase.Register(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) Login(ctx context.Context, params *models.LoginRequest) (*models.LoginResponse, error) {
	res, err := h.useCase.Login(ctx, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *handler) Logout(ctx context.Context, params *user.PostLogoutParams) error {
	err := h.useCase.Logout(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
