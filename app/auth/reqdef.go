package auth

import "github.com/golang-jwt/jwt"

type LoginDto struct {
	Nickname string `json:"nickname" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=2"`
}

type LoginAck struct {
	Token string `json:"token" validate:"required"`
}

type JWTClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
