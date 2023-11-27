package api

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
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

func CreateToken(username string , role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"role": role,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


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


func FindByName(username string) (*utils.Credentials, error) {
	var user utils.Credentials 
	db:=database.OpenDB()
	defer db.Close()
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil // User not found
	} else if err != nil {
		return nil, err // Other error
	}
	return &user, nil
}

func GetPassword(username string) (string, error) {
	var hashedPassword string
	db:=database.OpenDB()
	defer db.Close()
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		log.Printf("Error retrieving hashed password: %v", err)
		return "", err
	}
	return hashedPassword, nil
}
func VerifyPassword(hash, password string) bool {
	

	if hash == password {
		return true
	} else {
		log.Printf("Error in verify hashed password:")
		return false
	}


}

func GetRole( name string) (string,error)  {
	var Role string


	db:=database.OpenDB()
	defer db.Close()

	err := db.QueryRow("SELECT role FROM users WHERE username = $1", name).Scan(&Role)
	if err==sql.ErrNoRows{
	 
	 return "", fmt.Errorf("no role for this name")
	}else if err!=nil{
		
      return "", fmt.Errorf("Error retrieving Role ")
	}


	return Role,nil



	
}

func CheckRole(role string) bool {
	
	if role == "staff" {
		return true
	} else {
		return false
	}
	
}