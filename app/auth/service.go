package auth

import (
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
	"todo/app/user"
	"todo/config"
)

type AuthService struct {
	userService user.UserService
	logger      *zap.SugaredLogger
	config      config.Config
}

func NewAuthService(userService user.UserService, logger *zap.SugaredLogger, config config.Config) *AuthService {
	return &AuthService{userService, logger, config}
}

func (authService *AuthService) Login(loginDto LoginDto) LoginAck {
	nickname := loginDto.Nickname
	password := loginDto.Password
	user := authService.userService.FindByNickname(nickname)
	if user == nil {
		panic("Nickname or Password is incorrect")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		panic("Nickname or Password is incorrect")
	}

	tokenDuration := time.Duration(authService.config.Jwt.Expires) * time.Minute
	expirationTime := time.Now().Add(tokenDuration)
	claims := JWTClaim{
		UserID: user.Id.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(authService.config.Jwt.Secret))

	if err != nil {
		panic("There is an error!")
	}

	return LoginAck{
		Token: tokenString,
	}
}
