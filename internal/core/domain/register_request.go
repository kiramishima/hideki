package domain

import (
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

type RegisterRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (u *RegisterRequest) Hash256Password(password string) string {
	buf := []byte(password)
	pwd := sha3.New256()
	pwd.Write(buf)
	return hex.EncodeToString(pwd.Sum(nil))
}

func (u *RegisterRequest) BcryptPassword(password string) (string, error) {
	buf := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(buf, bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

func (u *RegisterRequest) ValidateBcryptPassword(password, password2 string) bool {
	byteHash := []byte(password)
	buf := []byte(password2)
	err := bcrypt.CompareHashAndPassword(byteHash, buf)
	if err != nil {
		return false
	}
	return true
}
