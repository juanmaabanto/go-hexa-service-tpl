package valueobject

import (
	"errors"
	"unicode"

	"github.com/juanmaabanto/go-hexa-service-tpl/utils/crypto"
)

const (
	minPasswordLength = 8
)

type Password string

func NewHashedPassword(plainPassword string) (Password, error) {
	if len(plainPassword) < minPasswordLength {
		return "", errors.New("password is too short")
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, char := range plainPassword {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return "", errors.New("complexity error")
	}

	hashed, err := crypto.HashPassword(plainPassword)
	if err != nil {
		return "", err
	}

	return Password(hashed), nil
}

func (p Password) Match(plainPassword string) bool {
	return crypto.CheckPasswordHash(plainPassword, p.String())
}

func (p Password) String() string {
	return string(p)
}
