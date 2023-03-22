package utils

func CreateUserValidator(username string, password string) (bool, string) {

	if len(username) < 0 {
		return false, "Invalid Username"
	}

	if len(password) < 6 {
		return false, "Invalid Password"
	}
	return true, "User validated successfully."
}
