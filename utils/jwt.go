package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKEy = os.Getenv("JWT_SECRET")

func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(2 * time.Hour).Unix()
	claims["̉email"] = email
	return token.SignedString([]byte(secretKEy))
}

func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKEy), nil
	})
	if err != nil {
		return "", err
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return "", errors.New("Invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Invalid token claims")
	}
	// email := claims["email"].(string)
	email := claims["email"].(string)
	return email, nil
}
