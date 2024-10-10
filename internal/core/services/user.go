package services

import "github.com/juanmaabanto/go-hexa-service-tpl/internal/core/domain/model"

type userService struct {
}

func (svc *userService) Create(
	email, firstName, lastName, password, alias, imageUrl string,
	active bool,
) (*model.User, error) {
	user, err := model.
		UserBuilder().
		Alias(alias).
		Email(email).
		FirstName(firstName).
		ImageURL(imageUrl).
		LastName(lastName).
		Password(password).
		Build()

	if err != nil {
		return nil, err
	}

	// here add logic

	return user, err
}
