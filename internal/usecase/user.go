package usecase

import (
	"context"
	"hulhay-mall/internal/apis/operations/user"
	"hulhay-mall/internal/models"
)

func (u *useCase) GetProfile(ctx context.Context, params *user.GetProfileParams) (*models.Users, error) {

	user, err := u.repo.GetProfile(ctx, params)
	if err != nil {
		return nil, err
	}

	res := &models.Users{
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		Username:  user.Username,
		UserPhoto: user.UserPhoto,
		CreatedAt: user.CreatedAt,
	}

	return res, nil
}
