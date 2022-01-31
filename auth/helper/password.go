package helper

import (
	"github.com/erizkiatama/rasalibrary-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", models.NewServerError("0100001", 500, err)
	}

	return string(hash), nil
}

func ComparePassword(password, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return models.NewClientError("0100002", "incorrect password", 400)
	}

	return nil
}
