---
sidebar_position: 7
---

# CLI

Go Gin Boilerplate 使用 [Cobra](https://github.com/spf13/cobra) 打造现代命令行工具，简化项目管理和操作。

## 介绍

我们可以把 Go Gin Boilerplate 看作是一个 CLI 工具，通过命令行工具，我们可以启动服务、创建新模块、查看版本信息等。

在入口文件 `main.go` 中，我们可以看到，当 Go Gin Boilerplate 启动时，会执行 `cmd.Execute()` 方法。

```go
func main() {
    cmd.Execute()
}
```

其中 `cmd` 是 `cmd/root.go` 中定义的 `rootCmd`。

在 `cmd/roots.go` 文件中，我们在这里定义 `Excute` 方法，用于执行命令，启动 cmd 工具。

```go
var rootCmd = &cobra.Command{
	Use:   "ggb",
	Short: "Go-Gin-Boilerplate is a development boilerplate based on the Gin framework, aimed at helping developers quickly build and develop web applications.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", "Welcome to Go-Gin-Boilerplate. Use -h to see more commands")
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

同时我们在 初始化函数 `init` 中，增加子命令，用于执行不同的操作。

```go
// You will additionally define flags and handle configuration in your init() function.
func init() {
	rootCmd.AddCommand(ServerStartCmd) // add server start command
	rootCmd.AddCommand(VersionCmd)     // add version command
	rootCmd.AddCommand(NewCmd)         // add new command
    // 。。。。add more commands here
}
```

程序运行到这，Go Gin Boilerplate 的 CLI 工具就已经初始化完成了，我们可以想象现在 Go Gin Boilerplate 就是一个 CLI，任何操作都通过命令实现，包括，启动服务、创建新模块、查看版本信息等。

目前 Go Gin Boilerplate 支持以下子命令：

- **init**：初始化项目
- **new**：创建新模块
- **server**：启动服务
- **version**：查看版本信息

## 安装

### 从源码安装

```sh
go install github.com/sanjayheaven/ggb@latest
```

## 子命令

### init

init 子命令用于初始化项目。

```sh
ggb init hello
```

这里 **hello** 为项目名称，执行命令之后，会在当前目录下创建 hello 项目。

### new

new 子命令用于创建新模块/项目。

**创建新模块**

在终端中执行以下命令，可以创建新模块

```sh
go run main.go new module hello
```

命令会依次生成以下文件：

- `internal/models/hello.go`
- `internal/services/hello.go`
- `internal/controllers/hello.go`
- `internal/router/hello.go`

各个文件依照的模板文件分别为：

- `web/template/model.tmpl`
- `web/template/service.tmpl`
- `web/template/controller.tmpl`
- `web/template/router.tmpl`

### server

server 子命令用于启动服务。

在终端中执行以下命令：

```sh
go run main.go server
```

将会启动 Go Gin Boilerplate 后端服务，默认端口 为 `8080`。

会依次完成以下操作：

1. 路由初始化
2. 日志工具初始化
3. 配置文件初始化
4. 服务优雅启动

在 `cmd/server.go` 文件中，我们可以看到，server 子命令的定义如下：

```go

func start() {

	// init router
	router.Init()
	r := router.Router

	// init logger
	logger.Init()
	logger := logger.LogrusLogger

	// load env config
	configs.Init()
	EnvConfig := configs.EnvConfig

	// connect database
	// mysql.Connect(&EnvConfig.Mysql)
	// connect redis
	// redis.Connect(&EnvConfig.Redis)

	// graceful shutdown
	server := &http.Server{
		Addr:    EnvConfig.Server.Port,
		Handler: r,
	}

	logger.Printf("👻 Server is now listening at port:  %s\n", EnvConfig.Server.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("server listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	i := <-quit
	logger.Println("server receive a signal: ", i.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown error: %s\n", err)
	}
	logger.Println("Server exiting")

}
```

> **值得注意的是：**
>
> - 经过慎重的考虑，在 server 启动的时候，我们不强制性加载数据库，包括 mysql 和 redis，因为有些时候，我们可能只是想要启动服务，而不需要连接数据库，这样可以节省一些资源。
> - 同样的，我们不主动在 各个包中，默认初始化，而且仅仅在 server run 的时候，才会初始化，这样可以保证各个命令互不干扰。

### version

version 子命令用于查看版本信息。

在终端中执行以下命令，可以查看相关的版本信息

```sh
go run main.go version
```

## 生成模版

我们采用了逆向工程的思路，用于生成新的模版。

这是因为我们考虑到，每个人的项目都不一样，所以我们不强制性的使用某种代码书写方式，而是采用了模版的方式，让用户自己定义模版。

通过 **example** 模块，我们来生成新的模版，也就是说，开发过程中，只要完成 **example** 模块的内容，就可以利用它生成新的模版代码。
相应的，这些模版可以用于生成新的模块。

主要思路是：

- 获取 **example** 模块的相关文件内容
- 替换相关内容为新的模块标量，例如 `example` 替换为 `{{.ModuleName}}`
- 生成新的模版文件

### 示例

在终端执行命令：

```sh
make gt
```

make gt 命令会执行 `scripts/gen-tmpl.go` 文件，该文件会执行 main 函数，

```go
func GenTmpl(moduleName string) error {

	tmplFiles := map[string]string{
		"web/template/router.tmpl":     tools.GetFile("internal/router/example.go"),
		"web/template/controller.tmpl": tools.GetFile("internal/controllers/example.go"),
		"web/template/service.tmpl":    tools.GetFile("internal/services/example.go"),
		"web/template/model.tmpl":      tools.GetFile("internal/models/example.go"),
	}

	for tmplPath, fileContent := range tmplFiles {
		err := createTmplByExampleModule(tmplPath, fileContent)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return err
		}
		fmt.Printf("Tmpl created: %s\n", tmplPath)
	}

	return nil
}

func main() {
	err := GenTmpl("example")
	if err != nil {
		fmt.Printf("Error generating template files: %v\n", err)
		return
	}
}
```

其中，`tmplFiles` 为需要生成的模版文件，`createTmplByExampleModule` 方法用于生成新的模版文件。

```go
// createTmplByExampleModule create template files for new module according to the existed example module
func createTmplByExampleModule(tmplPath, exampleFileContent string) error {

	moduleNameUpperFirst := "Example"
	moduleNamePlural := "examples"
	moduleName := "example"

	replacements := map[string]string{
		moduleNameUpperFirst: "{{.ModuleNameUpperFirst}}",
		moduleNamePlural:     "{{.ModuleNamePlural}}",
		moduleName:           "{{.ModuleName}}",
	}

	modifiedContent := replaceStrings(string(exampleFileContent), replacements)

	err := os.WriteFile(tmplPath, []byte(modifiedContent), 0644)
	if err != nil {
		return err
	}
	return nil

}
```

replaceStrings 方法用于替换相关内容为新的模块标量，例如 `example` 替换为 `{{.ModuleName}}`。

```go
// replaceStrings replace strings in input string according to the replacements map
func replaceStrings(input string, replacements map[string]string) string {
	for oldStr, newStr := range replacements {
		re := regexp.MustCompile(oldStr)
		input = re.ReplaceAllString(input, newStr)
	}
	return input
}
```

## 自定义命令

如果想要自定义命令，可以在 `cmd` 目录下创建新的命令文件，例如 `cmd/hello.go`。

```go
touch cmd/hello.go
```

在 `hello.go` 中，我们可以定义新的命令，例如 `HelloCmd`。

### 创建命令

```go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var HelloCmd = &cobra.Command{
    Use:   "hello",
    Short: "Say hello to the world",
    Long:  `Say hello to the world`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello World!")
    },
}
```

### 加入 rootCmd

确保自定义的命令创建完成之后，在 `cmd/root.go`的 init 中加入自定义的命令。

```go
// You will additionally define flags and handle configuration in your init() function.
func init() {
	rootCmd.AddCommand(HelloCmd) // add hello command
}
```
