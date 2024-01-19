package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// ErrorHandle is a middleware to handle panic
func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				// log error
				log.Println(err)

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

}
