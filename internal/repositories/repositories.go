package repositories

import (
	"context"
	"hulhay-mall/internal/apis/operations/user"
	"hulhay-mall/internal/models"

	"gorm.io/gorm"
)

type repositories struct {
	qry *gorm.DB
}

type Repositories interface {
	CreateStore(ctx context.Context, params *models.StoresRequest) error
	GetStores(ctx context.Context) ([]*models.Stores, error)
	GetStoreByID(ctx context.Context, storeID string) (*models.Stores, error)
	DeleteStoreByID(ctx context.Context, storeID string) error
	UpdateStoreByID(ctx context.Context, params *models.StoresRequest, storeID string) error

	Register(ctx context.Context, params *models.RegisterRequest) error
	GetByEmail(ctx context.Context, email string) (*models.Users, error)
	GetByUsername(ctx context.Context, username string) (*models.Users, error)
	GetByIdentifier(ctx context.Context, identifier string) (*models.Users, error)
	Login(ctx context.Context, params *models.LoginRequest) error
	Logout(ctx context.Context, params *user.PostLogoutParams) error
	GetProfile(ctx context.Context, params *user.GetProfileParams) (*models.Users, error)
}

func NewRepositories(q *gorm.DB) Repositories {
	return &repositories{qry: q}
}
