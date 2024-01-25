---
sidebar_position: 2
---

# 控制器

控制器是一个处理程序，它处理传入的 HTTP 请求，并返回响应。每个控制器都是一个 `gin` 处理程序。

## 介绍

Go Gin Boilerplate 中的控制器 位于 `internal/controller` 目录下，我们可以看到 `public.go`、`user.go`、`example.go` 等文件，这些文件就是我们的控制模块。

在控制器中，我们主要需要完成以下操作：

- 解析上下文，进行参数校验
- 调用服务层，进行业务处理
- 返回响应

## 解析上下文，参数校验

### 解析上下文

在 `gin` 框架中，我们可以通过 `c` 参数获取请求参数。

```go
func GetHello(c *gin.Context) {
    // 获取请求参数
    name := c.Query("name")
    // ...
}
```

<!-- ### 参数校验 -->

## 服务调用

在控制器中，我们可以调用服务层，进行业务处理。

```go

type ExampleController struct{}

func (exampleController *ExampleController) GetExample(ctx *gin.Context) {
	exampleIdStr := ctx.Query("exampleId")
	exampleId, err := strconv.Atoi(exampleIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error param"})
		return
	}

	res := exampleService.GetExample(exampleId)
	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Not Found"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
```

在上面的代码中，GetExample 方法完成了以下几个内容

- 解析请求参数，获取 `exampleId`
- 调用 `exampleService.GetExample` 方法，获取 `example` 数据
- 返回响应
