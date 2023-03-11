package repository

import (
	"errors"
	"server/api/helper"
	"server/api/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(u domain.User) domain.User
	VerifyCredential(username, password string) (domain.User, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &userRepository{
		connection: connection,
	}
}

func (r *userRepository) Create(u domain.User) domain.User {
	u.Password = helper.HashAndSalt([]byte(u.Password))
	r.connection.Create(&u)
	return u
}

func (r *userRepository) VerifyCredential(username, password string) (domain.User, error) {
	var user domain.User

	res := helper.ComparedPassword(user.Password, []byte(password))
	if !res {
		return user, errors.New("password tidak sama")
	}

	r.connection.Find(&user, "username = ? AND password = ?", username, password)

	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	return user, nil
}
