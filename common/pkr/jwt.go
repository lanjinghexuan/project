package pkr

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const jwtSecret = "AllYourBase"

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

// GetToken 生成 JWT Token
func GetToken(UserId string) (string, error) {
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        UserId,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return ss, nil
}

// ParseToken 解析 JWT Token
func ParseToken(tokenString string) (string, error) {
	var claims MyCustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	}, jwt.WithLeeway(time.Second))
	if err != nil {
		return "", fmt.Errorf("parse token error: %v", err)
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}
	return claims.ID, nil
}
