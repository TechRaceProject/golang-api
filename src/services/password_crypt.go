package services

import "golang.org/x/crypto/bcrypt"

const DefaultCost = 10

func HashPassword(password string) ([]byte, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)

	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}
