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

	if params.Role == `` {
		params.Role = shared.CUSTOMER
	}

	if params.Role != `` {
		if !shared.IsAllowedAccountRole[params.Role] {
			return errors.New("invalid role")
		}
	}

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
	uniqueIDRole := fmt.Sprintf("%sas%s", uniqueID, params.Role)

	encryptedPassword, err = shared.EncryptPassword(*params.Password, uniqueIDRole)
	if err != nil {
		return err
	}

	fullName := fmt.Sprintf("%s %s", *params.FirstName, params.LastName)

	err = u.repo.Register(ctx, &models.RegisterRequest{
		Email:    params.Email,
		FullName: fullName,
		UniqueID: uniqueIDRole,
		Username: params.Username,
		Password: &encryptedPassword,
		Role:     params.Role,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *useCase) Login(ctx context.Context, params *models.LoginRequest) (*models.LoginResponse, error) {

	var (
		err  error
		user *models.Users
	)

	user, err = u.repo.GetByUsername(ctx, *params.Username)

	// Check username
	if err != nil && user == nil {
		return nil, errors.New("username not found")
	}

	// Check password
	passUniqueID := fmt.Sprintf("%s%s", *params.Password, user.UniqueID)
	err = shared.CheckPassword(passUniqueID, user.Password)
	fmt.Println("[DEBUG] ---- err :", err)
	if err != nil {
		return nil, errors.New("wrong password")
	}

	// Check isLogin
	if user.IsLogin == true {
		return nil, errors.New("you have logged in")
	}

	err = u.repo.Login(ctx, params)
	if err != nil {
		return nil, err
	}

	res := &models.LoginResponse{}
	res.FullName = user.FullName
	res.Username = user.Username
	res.Identifier = user.UniqueID

	return res, nil
}

func (u *useCase) Logout(ctx context.Context, params *user.PostLogoutParams) error {

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

func (u *useCase) GetProfile(ctx context.Context, params *user.GetProfileParams) (*models.Users, error) {

	user, err := u.repo.GetProfile(ctx, params)
	if err != nil {
		return nil, err
	}

	res := &models.Users{
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		Username:  user.Username,
		UserPhoto: user.UserPhoto,
		CreatedAt: user.CreatedAt,
	}

	return res, nil
}
