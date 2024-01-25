---
sidebar_position: 2.3
---

# 中间件

中间件是一个函数，它可以访问请求和响应对象，以及应用程序的请求/响应循环中的下一个中间件函数。下一个中间件函数通常由名为 next 的变量表示。

## 介绍

在目录 `internal/middlewares` 中，我们用于存放相关的中间件，并且定义了一个名为 `middleware.go` 的文件，作为中间件包的入口。

在包入口文件中，我们完成初始化中间件的工作，然后将其导出，以便在应用程序中使用。

同时在 包入口文件中，我们常加载全局中间件，例如：错误处理中间件、跨域中间件等。

```go
package router

import (
	"github.com/sanjayheaven/ggb/internal/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/sanjayheaven/ggb/api/swagger" // docs.go
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	// public routes, no auth required
	LoadPublicRoutes(Router)

	// Global middlewares
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())

	// user routes
	LoadUserRoutes(Router)

	// example routes
	LoadExampleRoutes(Router)

	// init swagger
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}


```

## 目前已有的中间件

- [JWT 鉴权中间件](https://github.com/sanjayheaven/go-gin-boilerplate/tree/main/internal/middlewares/jwt.go)
- [跨域中间件](https://github.com/sanjayheaven/go-gin-boilerplate/blob/main/internal/middlewares/cors.go)
- [错误处理中间件](https://github.com/sanjayheaven/go-gin-boilerplate/blob/main/internal/middlewares/errorHandle.go)

### JWT 鉴权中间件

JWT 鉴权中间件，用于验证请求头中的 `Authorization` 字段是否合法，如果合法则将用户信息写入上下文中，否则返回错误信息。

```go
// Jwt middleware
func Jwt() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}
		parts := strings.Split(tokenStr, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}

		token := parts[1]
		claims, err := utils.JwtVerify(token)

		if err != nil || claims == nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}
		ctx.Next()

	}
}
```

### 跨域中间件

跨域中间件，用于解决跨域问题，允许所有的请求来源。

目前 Go-Gin-Boilerplate 脚手架使用的是 [gin-contrib/cors](https://github.com/gin-contrib/cors) 中间件。

```go
// Cors middleware
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}
```

### 错误处理中间件

错误处理中间件，用于处理请求过程中的错误，例如：panic、异常等。

```go
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

```

## 如何使用中间件

在 Gin 中，中间件是通过 gin.HandlerFunc() 函数签名表示的。因此，我们可以将我们自己的函数转换为 gin.HandlerFunc()，并将其作为参数传递给 gin.Engine.Use() 方法。

### 全局中间件

全局中间件将在每个请求上运行，因此它们不应该执行任何特定于请求的操作。例如，如果我们想要记录每个请求的信息，我们可以使用全局中间件。

默认在 `middleware.go` 中间件包入口中 加载全局中间件，例如：加载错误处理中间件、跨域中间件等。

```go
func init() {
    // ...

	// Global middlewares
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())

    // ...
}
```

### 路由中间件

路由中间件将在特定路由上运行，因此它们可以执行特定于路由的操作。例如，如果我们想要针对 Swagger 文档生成器加载权限中间件，我们可以使用路由中间件。

```go
func init() {
    // ...

    // init swagger
    Router.GET("/swagger/*any", middlewares.Jwt(), ginSwagger.WrapHandler(swaggerfiles.Handler))

    // ...
}

```

## 自定义中间件

我们可以通过实现 gin.HandlerFunc() 函数签名来创建自定义中间件。例如，我们可以创建一个名为 `logger.go` 的自定义中间件，用于记录每个请求的信息。

```go
// Logger middleware
func Logger() gin.HandlerFunc {

    return func(c *gin.Context) {
        // Start timer
        start := time.Now()

        // Process request
        c.Next()

        // Stop timer
        end := time.Now()
        latency := end.Sub(start)

        // Get request status code
        status := c.Writer.Status()

        log.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %#v\n%s",
            end.Format("2006/01/02 - 15:04:05"),
            status,
            latency,
            c.ClientIP(),
            c.Request.Method,
            c.Request.URL,
            c.Errors.String(),
        )
    }
}
```

## 更多

更多 Gin 中间件参考

- [gin-contrib](https://github.com/orgs/gin-contrib/repositories?q=middleware&type=all&language=&sort=)
