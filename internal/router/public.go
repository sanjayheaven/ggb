package routes

import (
	"go-gin-boilerplate/internal/controllers"

	"github.com/gin-gonic/gin"
)

func LoadPublicRoutes(r *gin.Engine) (*gin.RouterGroup, error) {

	publicController := new(controllers.PublicController)
	public := r.Group("/public")
	{
		public.GET("/ping", publicController.Ping)
	}
	return public, nil
}
