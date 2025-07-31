package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("RAHASIA_LO_JANGAN_SEBARIN")

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(1 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		// Make sure we're using HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("could not parse claims")
}
