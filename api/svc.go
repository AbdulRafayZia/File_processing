package api

// LoginService to provide user login with JWT token support
import (
	"fmt"
	"strings"
	"time"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte("secret -key")
)

type MyClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//	func verifyToken(tokenString string) error {
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			return secretKey, nil
//		})
//		if err != nil {
//			return err
//		}
//		if !token.Valid {
//			return fmt.Errorf("Invalid token")
//		}
//		return nil
//	}
func verifyToken(tokenString string) (*MyClaims, error) {
	tokenString = strings.TrimSpace(tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Provide the key used to sign the token
		return secretKey, nil
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
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		fmt.Println("Error extracting claims")
		return nil, err
	}

	return claims, nil
}
