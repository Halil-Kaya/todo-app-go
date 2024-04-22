package user

import (
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

	nicknameUser := service.FindByNickname(dto.Nickname)
	if nicknameUser != nil {
		panic("This nickname is already taken")
	}

	dto.Password, _ = HashPassword(dto.Password)

	user, err := service.userRepository.CreateUser(dto)

	if err != nil {
		service.logger.Error(err)
		return nil, err
	}
	return user, nil
}

func (service *UserService) FindByNickname(nickname string) *model.User {
	user := service.userRepository.FindByNickname(nickname)
	return user
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
