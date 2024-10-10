package model

import (
	"errors"
	"strings"

	"github.com/juanmaabanto/go-hexa-service-tpl/internal/core/domain/enum"
	"github.com/juanmaabanto/go-hexa-service-tpl/internal/core/domain/valueobject"
)

type userBuilder struct {
	user *User
	err  error
}

func UserBuilder() *userBuilder {
	return &userBuilder{user: &User{}}
}

func (b *userBuilder) Build() (*User, error) {
	if b.err != nil {
		return nil, b.err
	}

	if b.user.email.String() == "" {
		return nil, errors.New("email can't be empty")
	}

	if b.user.password.String() == "" {
		return nil, errors.New("password can't be empty")
	}

	return b.user, nil
}

func (b *userBuilder) Status(status enum.UserStatus) *userBuilder {
	if b.err == nil {
		b.user.status = status
	}
	return b
}

func (b *userBuilder) Alias(alias string) *userBuilder {
	if b.err == nil {
		b.user.alias = strings.TrimSpace(alias)
	}
	return b
}

func (b *userBuilder) Email(email string) *userBuilder {
	if b.err == nil {
		if strings.TrimSpace(email) == "" {
			b.err = errors.New("email empty")
			return b
		}

		b.user.email, b.err = valueobject.NewEmail(email)
		if b.err != nil {
			b.err = errors.New("field validation error")
			return b
		}
	}
	return b
}

func (b *userBuilder) FirstName(firstName string) *userBuilder {
	if b.err == nil {
		b.user.firstName = strings.TrimSpace(firstName)
	}
	return b
}

func (b *userBuilder) LastName(lastName string) *userBuilder {
	if b.err == nil {
		b.user.lastName = strings.TrimSpace(lastName)
	}
	return b
}

func (b *userBuilder) Password(plainPassword string) *userBuilder {
	if b.err == nil {
		if strings.TrimSpace(plainPassword) == "" {
			b.err = errors.New("password empty")
			return b
		}

		b.user.password, b.err = valueobject.NewHashedPassword(plainPassword)
		if b.err != nil {
			b.err = errors.New("field validation error")
			return b
		}
	}
	return b
}

func (b *userBuilder) ImageURL(imageURL string) *userBuilder {
	if b.err == nil && strings.TrimSpace(imageURL) != "" {
		b.user.imageURL, b.err = valueobject.NewURL(imageURL)
		if b.err != nil {
			b.err = errors.New("field validation error")
		}
	}
	return b
}
