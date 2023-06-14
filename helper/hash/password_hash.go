package hash

import (
	"golang.org/x/crypto/bcrypt"
	"nandes007/blog-post-rest-api/helper"
)

func PasswordHash(password string) string {
	bytePassword := []byte(password)
	hashPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		helper.PanicIfError(err)
	}

	return string(hashPassword)
}
