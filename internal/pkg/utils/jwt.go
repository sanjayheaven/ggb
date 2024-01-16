package utils

import (
	"fmt"
	"go-gin-boilerplate/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims

	Uid      uint   `json:"uid"`
	Username string `json:"username"`
}

// generate tokens used for auth
func GenerateToken(claims *Claims) string {
	Envconfig := configs.EnvConfig

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 30)) // set expire time

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Envconfig.Jwt.Secret))
	if err != nil {
		panic(err)
	}
	return token
}

// verify token
func JwtVerify(tokenStr string) (*Claims, error) {

	Envconfig := configs.EnvConfig
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Envconfig.Jwt.Secret), nil
	})

	if !token.Valid || err != nil {
		return nil, fmt.Errorf("token invalid")
	}
	claims, ok := token.Claims.(*Claims)

	if float64(claims.ExpiresAt.Unix()) < float64(time.Now().Unix()) {
		return nil, fmt.Errorf("token expired")
	}

	if !ok {
		return nil, err
	}
	return claims, err

}
