package helpers

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// DefaultSalt represents the default cost value to be used if BCRYPT_SALT is not set or invalid
const DefaultSalt = bcrypt.DefaultCost

// HashPassword hashes the given password using bcrypt with a specified salt cost from environment variable or default.
func HashPassword(password string) (string, error) {
	salt := os.Getenv("BCRYPT_SALT")
	saltInt, err := strconv.Atoi(salt)
	if err != nil || saltInt < bcrypt.MinCost || saltInt > bcrypt.MaxCost {
		saltInt = DefaultSalt
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), saltInt)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
