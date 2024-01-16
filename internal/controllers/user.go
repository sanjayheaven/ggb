package controllers

import (
	"go-gin-boilerplate/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

type UserController struct{}

type LoginByUsernamePasswordRequest struct {
	Usernmae string `json:"username" default:"admin"`
	Password string `json:"password" default:"123456"`
}

// @Router /users/loginByUsernamePassword [post]
// @Description Login By Username Password
// @Tags User
// @Param data body LoginByUsernamePasswordRequest true "username、password"
func (userController *UserController) LoginByUsernamePassword(ctx *gin.Context) {

	data := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	username := data["username"].(string)
	password := data["password"].(string)

	if username == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error param"})
		return
	}

	token := userService.LoginByUsernamePassword(username, password)
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Username or Password Error"})
		return
	}

	ctx.String(http.StatusOK, token)

}
