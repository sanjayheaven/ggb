package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// ErrorHandle is a middleware to handle panic
func ErrorHandle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// log error

			// print stack trace
			debug.PrintStack()

			// return error response
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})

			// abort request
			c.Abort()
		}
	}()
	c.Next()
}
