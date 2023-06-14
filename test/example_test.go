package test

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
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

func TestJwtTokenExampleParsingError(t *testing.T) {
	var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("That's not even a token")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		// Invalid signature
		fmt.Println("Invalid signature")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}
