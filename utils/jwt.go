package utils

import (
	"errors"
	"github.com/dcyar/fiber-books-api/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateJwtToken(email string) (string, error) {
	tokenHash := jwt.New(jwt.SigningMethodHS256)

	claims := tokenHash.Claims.(jwt.MapClaims)
	claims["identity"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token, err := tokenHash.SignedString([]byte(config.Config("JWT_SECRET")))

	if err != nil {
		return "", errors.New("Token generate error")
	}

	return token, nil
}
