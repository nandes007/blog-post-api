package test

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
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

func TestCreateNewSymmetricToken(t *testing.T) {
	token := "Q9ruLqUoIBFDpt3kxBCnvMf93-dSgobUOI76O358OPvEpWOuziztWXESCAc_bD5QE3AcI8g6DmzSmhvE-1nkuQ=="
	tokenByte := []byte(token)
	var (
		key []byte
		jt  *jwt.Token
	)

	key = tokenByte
	jt = jwt.New(jwt.SigningMethodHS256)
	s, err := jt.SignedString(key)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
}

func TestGenerateRandomString(t *testing.T) {
	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	randomString := base64.URLEncoding.EncodeToString(bytes)
	fmt.Println(randomString)
}

// Using this method when user login.
func TestBuildingEndSigningToken(t *testing.T) {
	hmacSampleSecret := "Q9ruLqUoIBFDpt3kxBCnvMf93-dSgobUOI76O358OPvEpWOuziztWXESCAc_bD5QE3AcI8g6DmzSmhvE-1nkuQ==" // Example HMAC secret key //self doc
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "asd",                                      // This is subject, you can bind email/userId // self doc
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // Expire token
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))

	fmt.Println(tokenString, err)
}

// Using this in middleware when user access a guarded route.
func TestParsingAndValidatingTokenUsingHMAC(t *testing.T) {
	hmacSampleSecret := "Q9ruLqUoIBFDpt3kxBCnvMf93-dSgobUOI76O358OPvEpWOuziztWXESCAc_bD5QE3AcI8g6DmzSmhvE-1nkuQ==" // Example HMAC secret key //self doc
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk0Mzc0MjUsInN1YiI6ImFzZCJ9.iZy7kyCPUlXeQLKZMB4k4tGH2dZY2ADA8cJOJvB9GGg"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSampleSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["sub"], claims["exp"])
	} else {
		fmt.Println(err)
	}
}

func TestReadEnvFile(t *testing.T) {
	file, err := os.Open("../.env")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Invalid line in .env file: %s", line)
			continue
		}
		key := parts[0]
		value := parts[1]
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading .env file")
	}

	// Access the credential as environment variable
	apiKey := os.Getenv("API_KEY")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	// Use the credential in code.
	log.Printf("API Key : %s", apiKey)
	log.Printf("Username : %s", username)
	log.Printf("Password : %s", password)
}
