package repositories

import (
	"context"
	"hulhay-mall/internal/models"

	"gorm.io/gorm"
)

type repositories struct {
	qry *gorm.DB
}

type Repositories interface {
	CreateStore(ctx context.Context, params models.StoresRequest) error
}

func NewRepositories(q *gorm.DB) Repositories {
	return &repositories{qry: q}
}
