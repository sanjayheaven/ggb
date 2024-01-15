package services

import (
	"go-gin-boilerplate/internal/models"
	"go-gin-boilerplate/internal/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct{}

// @name LoginByUsernamePassword
// @description LoginByUsernamePassword
// @return string
func (userService *UserService) LoginByUsernamePassword(username string, password string) string {

	user := models.User{
		Username: username,
	}
	res := models.DB.First(&user, "username = ?", username)
	if res.Error != nil || res.RowsAffected == 0 {
		return ""
	}

	token := utils.GenerateToken(jwt.MapClaims{
		"uid":      user.ID,
		"username": "admin",
	})

	return token
}

// 注册
func (userService *UserService) Register() {

}

// 生成token
