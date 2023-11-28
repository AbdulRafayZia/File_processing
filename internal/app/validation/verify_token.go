package validation

import (
	"fmt"
	"strings"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(tokenString string) (*utils.MyClaims, error) {
	tokenString = strings.TrimSpace(tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &utils.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Provide the key used to sign the token
		return utils.SecretKey, nil
	})

	// Check for errors
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("Invalid signature")
		} else {
			fmt.Println("Error parsing token:", err)
		}
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		fmt.Println("Invalid token")
		return nil, err
	}

	// Extract claims
	claims, ok := token.Claims.(*utils.MyClaims)
	if !ok {
		fmt.Println("Error extracting claims")
		return nil, err
	}

	return claims, nil
}