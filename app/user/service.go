package user

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"todo/app/exception"
	"todo/app/model"
)

type UserService interface {
	FindByNickname(nickname string) *model.User
	FindById(id string) *model.User
	CreateUser(dto UserCreateDto) (*model.User, error)
}

type userService struct {
	userRepository UserRepository
	logger         *zap.SugaredLogger
}

func NewUserService(userRepository UserRepository, logger *zap.SugaredLogger) UserService {
	return &userService{userRepository, logger}
}

func (service *userService) CreateUser(dto UserCreateDto) (*model.User, error) {

	nicknameUser := service.FindByNickname(dto.Nickname)
	if nicknameUser != nil {
		return nil, exception.NewNicknameIsAlreadyTaken()
	}

	dto.Password, _ = HashPassword(dto.Password)

	user, err := service.userRepository.CreateUser(dto)

	if err != nil {
		service.logger.Error(err)
		return nil, err
	}
	return user, nil
}

func (service *userService) FindByNickname(nickname string) *model.User {
	user := service.userRepository.FindByNickname(nickname)
	return user
}

func (service *userService) FindById(userId string) *model.User {
	user := service.userRepository.FindById(userId)
	return user
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
