package routes

import (
	"go-gin-boilerplate/internal/controllers"

	"github.com/gin-gonic/gin"
)

var publicController = new(controllers.PublicController)

func LoadPublicRoutes(r *gin.Engine) *gin.RouterGroup {

	public := r.Group("/public")
	{
		public.GET("/ping", publicController.Ping)
	}
	return public
}
