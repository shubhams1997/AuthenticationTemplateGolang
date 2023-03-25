package helper

import (
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, salt string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateRandomString() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return os.Getenv("DEFAULT_SALT")
	}
	return uuid.String()
}

func CheckPassword(password string, hash string, salt string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
}
