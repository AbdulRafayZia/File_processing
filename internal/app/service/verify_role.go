package service


func CheckRole(role string) bool {
	
	if role == "staff" {
		return true
	} else {
		return false
	}
	
}