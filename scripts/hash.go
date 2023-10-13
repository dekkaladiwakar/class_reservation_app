package scripts

import "golang.org/x/crypto/bcrypt"

func ConvertStringToHash(s string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(byte), nil
}

func CompareHashAndPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
