---
sidebar_position: 1
---

# 路由

路由是一个 URL 到一个特定的处理程序的映射。对于任何给定的请求，路由器将检查请求的路径与定义的路由列表进行匹配，如果找到匹配项，它将执行与该路由关联的处理程序。

## 介绍

在 Go Gin Boilerplate 中，路由是通过 `gin` 框架实现的，我们可以在 `internal/router` 目录下看到 `router.go` 文件，这个文件就是我们的路由入口文件。

路由按照模块区分。我们可以在 `internal/router` 目录下看到 `public.go`、`user.go`、`example.go` 等文件，这些文件就是我们的模块。

在入口文件中，我们完成 路由初始化、加载中间件、路由注册等操作。

```go
func Init() {
	Router = gin.Default()

	// Global middlewares
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())

	// public routes, no auth required
	LoadPublicRoutes(Router)

	// user routes
	LoadUserRoutes(Router)

	// example routes
	LoadExampleRoutes(Router)

	// init swagger
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
```

## 已有路由

在各个模块中，提供加载函数，用于 在 `router.go` 文件中加载路由。  
例如：我们可以看到 `internal/router/public.go` 文件，这个文件就是我们的公共模块路由文件。

```go
func LoadPublicRoutes(r *gin.Engine) *gin.RouterGroup {

    public := r.Group("/public")
    {
        public.GET("/getHello", publicController.GetHello)
    }
    return public
}
```

## 路由注册

路由注册是指将路由与处理程序进行绑定，当请求到来时，路由器会根据请求的路径，找到对应的处理程序进行处理。

例如在 `router.go` 文件中，我们可以看到 `LoadPublicRoutes` 函数，这个函数用于注册公共模块的路由。

### 注册新路由

- 创建新路由模块文件，例如 `internal/router/hello.go`。
- 在 `hello.go` 文件中，创建 加载路由的函数。

```go

func LoadHelloRoutes(r *gin.Engine) *gin.RouterGroup {

	hello := r.Group("/hellos")
	{
		hello.GET("/ping", helloController.Ping)
	}
	return hello
}

```

其中 `helloController.Ping` 是一个控制器函数。
我们可以在 `internal/controller` 目录下创建 `hello.go` 文件，这个文件就是我们的控制器文件。

- 在 `router.go` 文件中，导入新路由模块文件。

```go
func Init() {
	Router = gin.Default()

    // ... other code

	// hello routes
	LoadHelloRoutes(Router)

    // ... other code

}
```
