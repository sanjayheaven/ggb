<h1 align="center" >Go Gin Boilerplate</h1>

<div align="center">

一个基于 Gin 框架的开发脚手架，旨在帮助开发者快速搭建和开发 Web 应用程序

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sanjayheaven/go-gin-boilerplate)
![CI](https://github.com/sanjayheaven/go-gin-boilerplate/workflows/Go/badge.svg)

[Swagger 接口](https://google.com)


</div>

## 关键词

[Go](https://github.com/golang/go)
[Gin](https://github.com/gin-gonic/gin)
[Viper](https://github.com/spf13/viper)
[Gorm](https://github.com/go-gorm/gorm)
[Gin-Swagger](https://github.com/swaggo/gin-swagger)
[Air](https://github.com/cosmtrek/air)
[Logrus](https://github.com/sirupsen/logrus)
[Lumberjack](https://github.com/natefinch/lumberjack)

## 特性

- **模块化架构**: 遵循 project-layout 规范, 结构清晰，便于扩展和维护。
- **强大的中间件支持**: 整合常用中间件，轻松实现日志、认证等功能。
- **数据库集成**: 提供简单的接口和 ORM 工具，方便与数据库交互。
- **日志记录**: 集成日志记录功能，方便跟踪应用程序运行状态。

## 快速开始

### 拉取项目

```sh
git clone https://github.com/sanjayheaven/go-gin-boilerplate.git
```

### 进入目录，安装依赖

```sh
cd go-gin-boilerplate
go mod download
```

### 设置配置文件

### 运行应用程序

- 使用 [air](https://github.com/cosmtrek/air) 运行项目 **【推荐】**

```sh
air
```

- 使用 [go run](https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program) 运行项目

```sh
go run cmd/main.go
```

## 部署
