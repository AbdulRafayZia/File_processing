
package validation
import (
	"net/http"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
)

func CheckUserValidity(w http.ResponseWriter, r *http.Request, request utils.Credentials)(bool, error){
	role, err := service.GetRole(request.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthozied username", http.StatusUnauthorized)
		return false , err

	}

	validRole := service.CheckUserRole(role)
	if !validRole {
		http.Error(w, "Not a user", http.StatusUnauthorized)
		return false , err


	}

	user, err := service.FindByName(request.Username)
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return false , err

	}
	hashedPassword, err := service.GetPassword(request.Username)
	if err != nil {
		http.Error(w, "error getting hashed password", http.StatusBadRequest)
		return false , err

	}

	// Verify the password
	validPassword := VerifyPassword(hashedPassword, request.Password)
	if !validPassword {

		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return false , err


	}
	if user == nil && !validPassword {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "invalid credentials ", http.StatusUnauthorized)

		return false , err

	}
	return true, nil


}
