package repositories

import (
	"context"
	"hulhay-mall/internal/apis/operations/user"
	"hulhay-mall/internal/models"
)

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
