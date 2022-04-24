package main

import (
	// "go/token"
	"errors"
	"fmt"
	"time"

	// "github.com/golang-jwt/jwt/v4"
	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var key = []byte("secret")

func GenToken(username string, password string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
			Issuer:    "zhangmeng",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	// 注意这个地方一定要是直接切片不能是字符串
	return token.SignedString(key)
}

func ParseToken(tokenstring string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenstring, &MyClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return key, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func main() {
	s, e := GenToken("zm", "123")
	if e != nil {
		panic(e)
	}
	fmt.Printf("s: %v\n", s)

	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InptIiwicGFzc3dvcmQiOiIxMjMiLCJleHAiOjE2NTAyODY5NDIsImlzcyI6InpoYW5nbWVuZyJ9.criVL8WTvkB8vXY8AUVzJhz7BGwKdx1xIxknAcYHQdo"

	token := s

	mc, err := ParseToken(token)
	if err != nil {
		panic(err)
	}

	fmt.Printf("mc.Username: %v\n", mc.Username)
	fmt.Printf("mc.Password: %v\n", mc.Password)

}
