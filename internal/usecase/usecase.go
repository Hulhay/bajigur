package usecase

import (
	"context"
	"hulhay-mall/internal/apis/operations/auth"
	"hulhay-mall/internal/apis/operations/user"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/repositories"
)

type useCase struct {
	repo repositories.Repositories
}

type UseCase interface {
	// Store UC
	CreateStore(ctx context.Context, params *models.StoresRequest) error
	GetStores(ctx context.Context) ([]*models.Stores, error)
	GetStoreByID(ctx context.Context, storeID string) (*models.Stores, error)
	DeleteStoreByID(ctx context.Context, storeID string) error
	UpdateStoreByID(ctx context.Context, params *models.StoresRequest, stringID string) error

	// Auth UC
	Register(ctx context.Context, params *models.RegisterRequest) error
	Login(ctx context.Context, params *models.LoginRequest) (*models.LoginResponse, error)
	Logout(ctx context.Context, params *auth.PostLogoutParams) error

	// User UC
	GetProfile(ctx context.Context, params *user.GetProfileParams) (*models.Users, error)
}

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
