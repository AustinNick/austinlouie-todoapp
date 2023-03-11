package service

import (
	"server/api/model/domain"
	"server/api/model/web"
	"server/api/repository"
)

type UserService interface {
	Register(b web.UserRegisterRequest) (domain.User, error)
	Login(b web.UserLoginRequest) (domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Register(b web.UserRegisterRequest) (domain.User, error) {
	var user domain.User
	user.Username = b.Username
	user.Password = b.Password

	user = s.userRepository.Create(user)
	return user, nil
}

func (s *userService) Login(b web.UserLoginRequest) (domain.User, error) {
	user, err := s.userRepository.VerifyCredential(b.Username, b.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
