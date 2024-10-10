package model

import (
	"github.com/juanmaabanto/go-hexa-service-tpl/internal/core/domain/enum"
	"github.com/juanmaabanto/go-hexa-service-tpl/internal/core/domain/valueobject"
)

type User struct {
	ID        string
	email     valueobject.Email
	password  valueobject.Password
	firstName string
	lastName  string
	alias     string
	imageURL  valueobject.URL
	status    enum.UserStatus
}

func (u User) GetEmail() valueobject.Email {
	return u.email
}

func (u User) GetPassword() valueobject.Password {
	return u.password
}

func (u User) GetFirstName() string {
	return u.firstName
}

func (u User) GetLastName() string {
	return u.lastName
}

func (u User) GetAlias() string {
	return u.alias
}

func (u User) GetImageURL() valueobject.URL {
	return u.imageURL
}

func (u User) GetStatus() enum.UserStatus {
	return u.status
}
