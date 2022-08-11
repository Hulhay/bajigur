package usecase

import (
	"context"
	"errors"
	"fmt"
	"hulhay-mall/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *useCase) Register(ctx context.Context, params *models.RegisterRequest) error {

	var err error

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

	passUUID := fmt.Sprintf("%s%s", *params.Password, uniqueID)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(passUUID), 14)
	hashedPasswordStr := string(hashedPassword)

	err = u.repo.Register(ctx, &models.RegisterRequest{
		Email:     params.Email,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		FullName:  fullName,
		UniqueID:  uniqueID,
		Username:  params.Username,
		Password:  &hashedPasswordStr,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *useCase) Login(ctx context.Context, params *models.LoginRequest) error {
	return nil
}
