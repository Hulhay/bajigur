package handlers

import (
	"context"
	"errors"
	"hulhay-mall/internal/apis/operations/user"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/shared"
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

func (h *handler) Logout(ctx context.Context, params *user.PatchLogoutParams) error {
	err := h.useCase.Logout(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) GetProfile(ctx context.Context, params *user.GetProfileParams) (*models.Users, error) {

	var (
		err error
		res *models.Users
	)

	_, err = shared.GetRole(params.Identifier)
	if err != nil {
		return nil, errors.New("invalid identifier")
	}

	res, err = h.useCase.GetProfile(ctx, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}
