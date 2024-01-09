package routes

import (
	"github.com/gin-gonic/gin"
)

func Init() (*gin.Engine, error) {
	r := gin.Default()

	// public routes, no auth required
	LoadPublicRoutes(r)


	// example routes
	

	return r, nil
}
