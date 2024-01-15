package routes

import (
	"go-gin-boilerplate/internal/controllers"

	"github.com/gin-gonic/gin"
)

var userController = new(controllers.UserController)

func LoadUserRoutes(r *gin.Engine) *gin.RouterGroup {

	user := r.Group("/users")
	{
		user.GET("/loginByUsernamePassword", userController.LoginByUsernamePassword)
	}
	return user
}
