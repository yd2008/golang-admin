package util

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
}

func CompareWithPassword(hasdPwd, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasdPwd), []byte(password))
}