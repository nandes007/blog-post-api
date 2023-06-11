package test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	now := time.Now()
	fmt.Println("Unix Format", now.Format("2006-01-02"))
}

func TestHashPassword(t *testing.T) {
	password := "password"
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	hashedPasswordString := string(hashPassword)

	fmt.Println(hashedPasswordString)
}

func TestCompareHashPassword(t *testing.T) {
	passwordInDB := "password"
	requestPassword := "password123"
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(passwordInDB), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	hashedPasswordString := string(hashPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPasswordString), []byte(requestPassword))

	if err != nil {
		fmt.Println("Password is missmatch")
	} else {
		fmt.Println("Password is match")
	}

}
