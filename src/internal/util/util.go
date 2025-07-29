package util

import "golang.org/x/crypto/bcrypt"

func EncriptedPassword(senha string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(hash), err
}
