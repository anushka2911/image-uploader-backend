package service

import (
	"errors"

	"github.com/anushkapandey/image_uploader_backend/model"
)

func ValidateCredentials(userDB map[string]*model.User, email, password string) (string, error) {
	user, exists := userDB[email]
	if !exists || user.IsDeleted {
		return "", errors.New("user does not exist")
	}

	if user.Password != password {
		return "", errors.New("invalid password")
	}
	return user.Role, nil
}

func UserExists(userDB map[string]*model.User, email string) (bool, error) {
	_, exists := userDB[email]
	return exists, nil
}

func AddUser(userDB map[string]*model.User, email, password, role string) error {

	userDB[email] = &model.User{
		Email:     email,
		Password:  password,
		Role:      role,
		IsDeleted: false,
	}
	return nil
}
