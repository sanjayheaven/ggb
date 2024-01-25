package services

import (
	"github.com/sanjayheaven/ggb/internal/models"
	"github.com/sanjayheaven/ggb/internal/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// @name LoginByUsernamePassword
// @description LoginByUsernamePassword
// @return string
func (userService *UserService) LoginByUsernamePassword(username string, password string) string {

	user := models.User{
		Username: username,
	}
	res := db.First(&user, "username = ?", username)
	if res.Error != nil || res.RowsAffected == 0 {
		return ""
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return ""
	}
	claims := utils.Claims{
		Username: user.Username,
		Uid:      user.ID,
	}
	token := utils.GenerateToken(&claims)

	return token
}
