package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/example/gin_framework/helper"
)

type JWTService interface {
	GenerateJWT(name string, duration time.Duration) (string, error)
	ValidateJWT(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	secretKey string
}

func NewJwtService() JWTService {
	return &jwtCustomClaims{secretKey: os.Getenv("SECRETKEY")}
}

func (jwtCustom *jwtCustomClaims) GenerateJWT(name string, duration time.Duration) (string, error) {
	payload, err := helper.NewPayload(name, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(jwtCustom.secretKey))
	return token, err
}

func (jwtCustom *jwtCustomClaims) ValidateJWT(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtCustom.secretKey), nil
	})
}
