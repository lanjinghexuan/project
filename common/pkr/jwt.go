package pkr

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

//代码从官方查询个人参考使用

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func GetToken(UserId string) (string, error) {
	//密钥
	mySigningKey := []byte("AllYourBase")

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //发布时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     //生效时间
			ID:        UserId,                                             //内容
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(ss, err)
	return ss, nil
}

func ParseToken(tokenString string) (string, error) {

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	}, jwt.WithLeeway(time.Second))
	if err != nil {
		fmt.Println(err)
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims.ID, nil
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
	return "", errors.New("unknown claims type, cannot proceed")
	// Output: bar test
}
