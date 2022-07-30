package repositories

import (
	"context"
	"fmt"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/shared"
	"time"
)

func (r *repositories) CreateStore(ctx context.Context, params models.StoresRequest) error {
	var stores models.Stores

	createdAt := time.Now()
	updatedAt := createdAt

	if params.StorePhoto == "" {
		params.StorePhoto = shared.DEFAULT_STORE_PHOTO
	}

	tx := r.qry.Begin()

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

	tx.Commit()

	fmt.Println("repo :", &stores)

	return nil
}
