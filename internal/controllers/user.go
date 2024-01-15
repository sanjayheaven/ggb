package controllers

import (
	"go-gin-boilerplate/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

type UserController struct{}

// @Router /users/loginByUsernamePassword [post]
// @Description Login By Username Password
// @Tags User
// @Param username body string true "username"
// @Param password body string true "password"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error"})
		return
	}
	ctx.String(http.StatusOK, token)

}
