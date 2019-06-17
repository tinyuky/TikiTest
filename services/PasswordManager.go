package services

type valicontent struct {
	message string
	check   bool
}

/*set up validate rule for check password*/
var validator = map[string]valicontent{
	`\w{6,}`:    {"The password must be at least 6 characters long", true},
	`\s`:        {"The password must not contain any whitespace", false},
	`[a-z]`:     {"The password must contain at least one uppercase and at least one lowercase letter", true},
	`[A-Z]`:     {"The password must contain at least one uppercase and at least one lowercase letter", true},
	`[@$!%*?&]`: {"The password must have at least one digit and symbol", true},
	`[0-9]`:     {"The password must have at least one digit and symbol", true},
}

/*VerifyPassword : funtion compare password with stored password*/
func VerifyPassword(inputpass string, storedpass []byte) (bool, error) {
	plaintext, err := decrypt(storedpass, "password")
	if err != nil {
		return false, err
	}
	if string(plaintext) != inputpass {
		return false, nil
	}
	return true, nil
}

/*ValidatePassword : function to validate password follow rules*/
func ValidatePassword(pass string) (bool, string, error) {
	msg, err := validate(pass)
	if err != nil || msg != "" {
		return false, msg, err
	}
	return true, "", nil
}

/*SetNewPassword : function encrypt new password*/
func SetNewPassword(pass string) ([]byte, error) {
	ciphertext, err := encrypt([]byte(pass), "password")
	return ciphertext, err
}

// /*GetPassword : function ge*/
// func GetPassword(pass []byte) (string, error) {
// 	plaintext, err := decrypt(pass, "password")
// 	return string(plaintext), err
// }

