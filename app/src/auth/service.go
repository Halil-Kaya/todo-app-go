package auth

import (
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
	"todo/app/core/exception"
	"todo/app/src/model"
	"todo/config"
)

type UserService interface {
	FindByNickname(nickname string) *model.User
	FindById(id string) *model.User
}

type AuthService struct {
	userService UserService
	logger      *zap.SugaredLogger
	config      config.Config
}

func NewAuthService(userService UserService, logger *zap.SugaredLogger, config config.Config) *AuthService {
	return &AuthService{userService, logger, config}
}

func (authService *AuthService) Login(loginDto LoginDto) (*LoginAck, error) {
	nickname := loginDto.Nickname
	password := loginDto.Password
	user := authService.userService.FindByNickname(nickname)
	if user == nil {
		return nil, exception.NewUnauthorized()
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, exception.NewUnauthorized()
	}

	tokenString := authService.CreateToken(user)

	return &LoginAck{
		Token: tokenString,
	}, nil
}

func (authService *AuthService) CreateToken(user *model.User) string {
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
		panic("There is an error while creating token!")
	}

	return tokenString
}

// returning userId
func (authService *AuthService) ValidateToken(token string) (string, error) {
	tokenClaim, err := jwt.ParseWithClaims(
		token,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(authService.config.Jwt.Secret), nil
		},
	)

	if err != nil {
		return "", exception.NewUnauthorized()
	}
	claims, ok := tokenClaim.Claims.(*JWTClaim)
	if !ok {
		panic("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		panic("token expired")
	}
	return claims.UserID, nil
}
