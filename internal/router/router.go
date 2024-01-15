package routes

import (
	"go-gin-boilerplate/internal/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-gin-boilerplate/api/swagger" // docs.go
)

func Init() (*gin.Engine, error) {
	r := gin.Default()

	// public := LoadPublicRoutes(r)
	// api := r.Group("/api")
	// {
	// 	public
	// }

	// public routes, no auth required
	LoadPublicRoutes(r)

	// user routes
	LoadUserRoutes(r)

	// Loading middlewares
	middlewares.LoadMiddlewares(r)

	// example routes
	LoadExampleRoutes(r)

	// init swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r, nil
}
