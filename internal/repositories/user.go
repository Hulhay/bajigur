package repositories

import (
	"context"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/shared"
	"time"
)

func (r *repositories) Register(ctx context.Context, params *models.RegisterRequest) error {
	var users *models.Users

	createdAt := time.Now()
	updatedAt := createdAt

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&users).Create(map[string]interface{}{
		"unique_id":  params.UniqueID,
		"username":   params.Username,
		"password":   params.Password,
		"full_name":  params.FullName,
		"email":      params.Email,
		"user_photo": shared.DEFAULT_USER_PHOTO,
		"created_at": createdAt,
		"updated_at": updatedAt,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *repositories) GetByEmail(ctx context.Context, email string) (*models.Users, error) {
	var users *models.Users

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&users).Where("email = ?", email).First(&users).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return users, nil
}

func (r *repositories) GetByUsername(ctx context.Context, username string) (*models.Users, error) {
	var users *models.Users

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&users).Where("username = ?", username).First(&users).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return users, nil
}
