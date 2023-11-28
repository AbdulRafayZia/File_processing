
package validation
import (
	"net/http"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
)

func CheckStaffValidity(w http.ResponseWriter, r *http.Request, request utils.Credentials)(bool, error){
	
	role, err := database.GetRole(request.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthozied username", http.StatusUnauthorized)

		return false , err
	}
	validRole := CheckStaffRole(role)
	if !validRole {
		http.Error(w, "Not a Staff Member", http.StatusUnauthorized)
		return false , err


	}

	user, err := database.FindByName(request.Username)
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return false , err

	}
	hashedPassword, err := database.GetPassword(request.Username)
	if err != nil {
		http.Error(w, "error in getting  from db", http.StatusBadRequest)
		return false , err

	}

	validPassword := VerifyPassword(hashedPassword, request.Password)
	if !validPassword {

		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "incorrect password ", http.StatusUnauthorized)
		return false , err


	}
	if user == nil && !validPassword && !validRole {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "invalid credentails ", http.StatusUnauthorized)

	}

 return true, nil
}
