package shared

import (
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

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordDB))
	if err != nil {
		return err
	}
	return nil
}
