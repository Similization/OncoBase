package utils

import (
	"errors"
	"fmt"
	"med/pkg/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

const tokenTTL = 5 * time.Minute

type tokenClaims struct {
	UserId   int    `json:"id"`
	UserRole string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(user model.User) (string, error) {
	expirationTime := time.Now().Add(tokenTTL)

	claims := &tokenClaims{
		UserId:   user.Id,
		UserRole: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ParseToken(accessToken string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Wrong method")
		}

		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("Token claims")
	}
	return claims, nil
}
