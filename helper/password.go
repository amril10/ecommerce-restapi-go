package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordHash), err
}

func VerifiyPassword(hashPassword string, password string) error {
	if hashPassword != password {
		return errors.New("invalid password")
	}

	return nil
	// err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	// return err
}
