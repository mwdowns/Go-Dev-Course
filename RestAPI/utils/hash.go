package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(email string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
