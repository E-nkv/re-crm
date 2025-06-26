package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtTTL = time.Hour * 24 * 7

var (
	JWT_SECRET string
)

func init() {
	JWT_SECRET = os.Getenv("JWT_SECRET")
}

func GenerateJWT(userID uint64, role string) (string, error) {
	expirationTime := time.Now().Add(jwtTTL).Unix()

	claims := jwt.MapClaims{
		"userID":   userID,
		"exp":      expirationTime,
		"userRole": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func DecodeJWT(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			return []byte(JWT_SECRET), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
