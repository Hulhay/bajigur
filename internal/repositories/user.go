package repositories

import (
	"context"
	"hulhay-mall/internal/apis/operations/user"
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

func (r *repositories) GetByIdentifier(ctx context.Context, identifier string) (*models.Users, error) {
	var users *models.Users

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&users).Where("unique_id = ?", identifier).First(&users).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return users, nil
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

func (r *repositories) Logout(ctx context.Context, params *user.PatchLogoutParams) error {
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

func (r *repositories) GetProfile(ctx context.Context, params *user.GetProfileParams) (*models.Users, error) {
	var users *models.Users

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&users).Where("unique_id = ?", params.Identifier).First(&users).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return users, nil
}
