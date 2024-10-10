package valueobject

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(email string) (Email, error) {
	if !validateEmailFormat(email) {
		return "", errors.New("invalid format")
	}

	return Email(email), nil
}

func (e Email) String() string {
	return string(e)
}

func validateEmailFormat(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return regex.MatchString(email)
}
