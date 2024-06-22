package util

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to has password: %w", err)
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	log.Println("HASHED PASS: ", hashedPassword)
	log.Println("PASS: ", password)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
