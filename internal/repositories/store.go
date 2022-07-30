package repositories

import (
	"context"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/shared"
	"time"
)

func (r *repositories) CreateStore(ctx context.Context, params models.StoresRequest) error {
	var stores *models.Stores

	createdAt := time.Now()
	updatedAt := createdAt

	if params.StorePhoto == "" {
		params.StorePhoto = shared.DEFAULT_STORE_PHOTO
	}

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&stores).Create(map[string]interface{}{
		"owner":         params.Owner,
		"store_address": params.StoreAddress,
		"store_name":    params.StoreName,
		"store_phone":   params.StorePhone,
		"store_photo":   params.StorePhoto,
		"created_at":    createdAt,
		"updated_at":    updatedAt,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *repositories) GetStores(ctx context.Context) ([]*models.Stores, error) {
	var stores []*models.Stores

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&stores).Find(&stores).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return stores, nil
}

func (r *repositories) GetStoreByID(ctx context.Context, storeID string) (*models.Stores, error) {
	var store *models.Stores

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&store).Where("id = ?", storeID).Find(&store).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return store, nil
}
