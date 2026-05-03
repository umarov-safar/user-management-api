package utils

import "golang.org/x/crypto/bcrypt"

const BcryptCost = 12

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		BcryptCost,
	)

	return string(hash), err
}

func VerifyPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	return err == nil
}
