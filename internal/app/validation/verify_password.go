package validation

import "log"

func VerifyPassword(hash, password string) bool {

	if hash == password {
		return true
	} else {
		log.Printf("Error in verify hashed password:")
		return false
	}

}
