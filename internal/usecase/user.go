package usecase

import (
	"context"
	"errors"
	"fmt"
	"hulhay-mall/internal/apis/operations/user"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/shared"

	"github.com/google/uuid"
)

func (u *useCase) Register(ctx context.Context, params *models.RegisterRequest) error {

	var (
		err               error
		encryptedPassword string
	)

	// Check the email
	_, err = u.repo.GetByEmail(ctx, params.Email.String())
	if err == nil {
		return errors.New("this email is already in use")
	}

	// Check the username
	_, err = u.repo.GetByUsername(ctx, *params.Username)
	if err == nil {
		return errors.New("this username is already in use")
	}

	uniqueID := uuid.New().String()

	fullName := fmt.Sprintf("%s %s", *params.FirstName, params.LastName)

	encryptedPassword, err = shared.EncryptPassword(*params.Password, uniqueID)
	if err != nil {
		return err
	}

	err = u.repo.Register(ctx, &models.RegisterRequest{
		Email:    params.Email,
		FullName: fullName,
		UniqueID: uniqueID,
		Username: params.Username,
		Password: &encryptedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *useCase) Login(ctx context.Context, params *models.LoginRequest) error {

	var (
		err  error
		user *models.Users
	)

	user, err = u.repo.GetByUsername(ctx, *params.Username)

	// Check username
	if err != nil && user == nil {
		return errors.New("username not found")
	}

	// Check password
	passUniqueID := fmt.Sprintf("%s%s", *params.Password, user.UniqueID)
	err = shared.CheckPassword(user.Password, passUniqueID)
	if err != nil {
		return errors.New("wrong password")
	}

	// Check isLogin
	if user.IsLogin == true {
		return errors.New("you have logged in")
	}

	err = u.repo.Login(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (u *useCase) Logout(ctx context.Context, params *user.PatchLogoutParams) error {

	var (
		err  error
		user *models.Users
	)

	user, err = u.repo.GetByIdentifier(ctx, params.Identifier)

	// Check username
	if err != nil && user == nil {
		return errors.New("username not found")
	}

	// Check is_login
	if user.IsLogin != true {
		return errors.New("you must login before logout")
	}

	err = u.repo.Logout(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
