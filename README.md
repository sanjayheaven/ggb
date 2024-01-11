# Go Gin Boilerplate

<div>

<img src="./assets/golang.png" align="right"/>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sanjayheaven/go-gin-boilerplate)
![CI](https://github.com/sanjayheaven/go-gin-boilerplate/workflows/Go/badge.svg)
![License](https://img.shields.io/github/license/sanjayheaven/go-gin-boilerplate)

[Swagger 接口](https://go-gin-boilerplate.gganbu.services/swagger/index.html)

</div>

## 项目介绍

一个基于 Gin 框架的开发脚手架，旨在帮助开发者快速搭建和开发 Web 应用程序。

- [English](./README_en.md)
- [中文介绍](./README.md)

## 关键词

[Go](https://github.com/golang/go)
[Gin](https://github.com/gin-gonic/gin)
[Viper](https://github.com/spf13/viper)
[Gorm](https://github.com/go-gorm/gorm)
[Gin-Swagger](https://github.com/swaggo/gin-swagger)
[Air](https://github.com/cosmtrek/air)
[Logrus](https://github.com/sirupsen/logrus)
[Lumberjack](https://github.com/natefinch/lumberjack)
[Jwt](https://github.com/golang-jwt/jwt)

## 特性

- **快速开发**: 使用 Gin 框架和相关工具，加速项目的开发和迭代过程。
- **简单易用**: 遵循 [project-layout](https://github.com/golang-standards/project-layout/tree/master) 规范, 提供清晰简单的代码结构，使新手也能轻松上手。
- **一体化日志系统**: 集成 Logrus 和 Lumberjack，实现全方位的日志记录和管理。
- **热重载**: 使用 Air 工具，支持热重载，提高开发效率。
- **数据库支持**: 集成 Gorm，支持主流数据库，如 MySQL、PostgreSQL 等。
- **灵活的中间件**: 整合常用中间件，轻松实现日志、认证、跨域、限流等功能。

## 快速开始

```sh
git clone https://github.com/sanjayheaven/go-gin-boilerplate.git
cd go-gin-boilerplate
go mod download
```

### 设置配置文件

- 进入 `configs` 目录，复制 `config.example.yaml` 文件并重命名为 `config.yaml`。

```sh
cp configs/config.example.yaml configs/config.yaml
```

- 修改 `config.yaml` 文件中的配置项。

```sh
vi configs/config.yaml
```

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

### docker-compose

使用 docker-compose 部署应用程序。  
确保 在服务器中已经安装 docker ，并且熟悉 docker compose 的使用

- 复制 [docker-compose.yml](./deployments/docker-compose.yml) 配置文件到 项目根目录 中
- 执行以下命令，启动应用程序

```sh
docker compose up -d
```

## 贡献

欢迎提交 [issue](https://github.com/sanjayheaven/go-gin-boilerplate/issues) 和 [PR](https://github.com/sanjayheaven/go-gin-boilerplate/pulls)，感谢你的支持。
