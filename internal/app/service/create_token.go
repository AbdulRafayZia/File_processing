package service

import (
	"time"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"role":     role,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(utils.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}