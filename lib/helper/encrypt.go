package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromString(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHashAndPassword(str1 string, str2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(str1), []byte(str2))
	if err != nil {
		return err
	}

	return nil
}
