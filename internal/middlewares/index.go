package middlewares

import "github.com/gin-gonic/gin"

func LoadMiddlewares(r *gin.Engine) *gin.Engine {

	r.Use(ErrorHandle)
	r.Use(Cors())

	return r
}
