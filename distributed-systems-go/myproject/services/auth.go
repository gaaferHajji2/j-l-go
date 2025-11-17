package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct{}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTOutput struct {
	Token     string    `json:"token"`
	ExpiresIn time.Time `json:"expires_in"`
}
