package user

import (
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"todo/app/model"
)

type UserService struct {
	userRepository UserRepository
	logger         *zap.SugaredLogger
}

func NewUserService(userRepository UserRepository, logger *zap.SugaredLogger) *UserService {
	return &UserService{userRepository, logger}
}

func (service *UserService) CreateUser(dto UserCreateDto) (*model.User, error) {
	dto.Password, _ = HashPassword(dto.Password)

	user, err := service.userRepository.CreateUser(dto)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
