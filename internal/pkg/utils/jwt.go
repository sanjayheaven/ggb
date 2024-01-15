package utils

import (
	configs "go-gin-boilerplate/configs"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generate tokens used for auth
func GenerateToken(data map[string]interface{}) string {
	Envconfig := configs.EnvConfig

	claims := jwt.MapClaims{}
	for k, v := range data {
		claims[k] = v
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(Envconfig.Jwt.Secret)
	if err != nil {
		panic(err)
	}
	return token
}

// VerifyToken verify token
func JwtVerify(tokenStr string) (jwt.MapClaims, error) {
	Envconfig := configs.EnvConfig

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// }
		return []byte(Envconfig.Jwt.Secret), nil
	})
	if !token.Valid || err != nil {
		panic("token invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, err
	} else {
		return nil, err
	}

}
