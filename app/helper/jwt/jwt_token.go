package jwt

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(subject int) (string, error) {
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
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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

func parseSubToInt(sub interface{}) (int, error) {
	if floatValue, ok := sub.(float64); ok {
		intValue := int(floatValue)
		return intValue, nil
	}

	return 0, fmt.Errorf("unable to parse 'sub' attribute as int")
}

func ParseUserToken(token string) (int, error) {
	hmacSecretKey := os.Getenv("JWT_SIGNING")

	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSecretKey), nil
	})

	if err != nil {
		return 0, errors.New("invalid token given")
	}

	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		sub, ok := claims["sub"]
		if ok {
			subInt, err := parseSubToInt(sub)
			if err != nil {
				log.Fatal(err)
			}
			return subInt, nil
		}
		return 0, errors.New("invalid token given")
	} else {
		return 0, errors.New("invalid token given")
	}
}

func ParseUserTokenV2(token string) (int, error) {
	tokenFormatted := strings.Replace(token, "Bearer ", "", 1)
	hmacSecretKey := os.Getenv("JWT_SIGNING")

	parseToken, err := jwt.Parse(tokenFormatted, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSecretKey), nil
	})

	if err != nil {
		return 0, errors.New("invalid token given")
	}

	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		sub, ok := claims["sub"]
		if ok {
			subInt, err := parseSubToInt(sub)
			if err != nil {
				log.Fatal(err)
			}
			return subInt, nil
		}
		return 0, errors.New("invalid token given")
	} else {
		return 0, errors.New("invalid token given")
	}
}
