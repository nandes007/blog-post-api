package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateToken(subject string) (string, error) {
	hmacSecretKey := os.Getenv("JWT_SIGNING")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSecretKey))

	return tokenString, err
}

func ValidateToken(token string) bool {
	hmacSecretKey := os.Getenv("JWT_SIGNING")

	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSecretKey), nil
	})

	if err != nil {
		return false
	}

	if _, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		return true
	} else {
		return false
	}
}
