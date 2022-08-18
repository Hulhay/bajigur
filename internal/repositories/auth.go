package repositories

import (
	"context"
	"hulhay-mall/internal/apis/operations/auth"
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
		"role":       params.Role,
		"user_photo": shared.DEFAULT_USER_PHOTO,
		"created_at": createdAt,
		"updated_at": updatedAt,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *repositories) Login(ctx context.Context, params *models.LoginRequest) error {
	var users *models.Users

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&users).Where("username = ?", params.Username).Updates(map[string]interface{}{
		"is_login":   true,
		"updated_at": time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *repositories) Logout(ctx context.Context, params *auth.PostLogoutParams) error {
	var users *models.Users

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&users).Where("unique_id = ?", params.Identifier).Updates(map[string]interface{}{
		"is_login":   false,
		"updated_at": time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
