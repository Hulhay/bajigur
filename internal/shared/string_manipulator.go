package shared

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password, uniqueID string) (string, error) {

	passUniqueID := fmt.Sprintf("%s%s", password, uniqueID)
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(passUniqueID), 14)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}

func CheckPassword(password, passwordDB string) error {

	err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GetRole(identifier string) (string, error) {

	if len(identifier) < 38 {
		return "", errors.New("invalid identifier")
	}

	role := string(identifier[38:])
	return role, nil
}
