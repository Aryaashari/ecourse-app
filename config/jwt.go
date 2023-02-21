package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("V2KuVnwF83zZLky")
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_EXPIRED = time.Duration(12) * time.Hour

type JWTClaims struct {
	jwt.StandardClaims
	Name  string `json:"name"`
	Email string `json:"email"`
}
