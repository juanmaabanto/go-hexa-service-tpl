package crypto

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	bytes, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword(bytes, []byte(password))
	return err == nil
}
