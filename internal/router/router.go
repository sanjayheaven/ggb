package routes

import (
	"go-gin-boilerplate/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Init() (*gin.Engine, error) {
	r := gin.Default()

	// public routes, no auth required
	LoadPublicRoutes(r)

	// Loading middlewares
	middlewares.LoadMiddlewares(r)

	// example routes
	LoadExampleRoutes(r)

	return r, nil
}
