package pkr

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// 生成JWT
func GenerateJWT(username int32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id": username,
	})
	jwtkey := make([]byte, 32)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtkey)

	return tokenString, err
}

// 解析JWT
func ParseJwt(tokenstring string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
	}
	if token.Valid {
		fmt.Println("Token is invalid")
	}
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Token is invalid;line 33")
	}
	
	return tokenClaims, nil
}
