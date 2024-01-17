---
sidebar_position: 1
---

# Go Gin Boilerplate

<div>

<img src="/img/golang.png" align="right"/>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sanjayheaven/go-gin-boilerplate)
![CI](https://github.com/sanjayheaven/go-gin-boilerplate/workflows/Go/badge.svg)
![License](https://img.shields.io/github/license/sanjayheaven/go-gin-boilerplate)

[Swagger 接口文档](https://go-gin-boilerplate.gganbu.services/swagger/index.html)

</div>

## 项目介绍

一个基于 Gin 框架的开发脚手架，旨在帮助开发者快速搭建和开发 Web 应用程序。

## 关键词

[Go](https://github.com/golang/go)
[Gin](https://github.com/gin-gonic/gin)
[Cobra](https://github.com/spf13/cobra)
[Viper](https://github.com/spf13/viper)
[Gorm](https://github.com/go-gorm/gorm)
[Gin-Swagger](https://github.com/swaggo/gin-swagger)
[Air](https://github.com/cosmtrek/air)
[Logrus](https://github.com/sirupsen/logrus)
[Lumberjack](https://github.com/natefinch/lumberjack)
[Jwt](https://github.com/golang-jwt/jwt)

## 特性

- **快速开发**: 使用 **Gin** 框架和相关工具，加速项目的开发和迭代过程。
- **简单易用**: 遵循 [project-layout](https://github.com/golang-standards/project-layout/tree/master) 规范, 提供清晰简单的代码结构，使新手也能轻松上手。
- **先进的 CLI 体验**: 使用 **Cobra** 打造现代命令行工具，简化项目管理和操作。
- **热重载**: 使用 **Air** 工具，支持热重载，提高开发效率。
- **一体化日志系统**: 集成 **Logrus** 和 **Lumberjack**, 实现全方位的日志记录和管理。
- **数据库支持**: 集成 **Gorm**, 支持主流数据库，如 MySQL、PostgreSQL 等。
- **灵活的中间件**: 整合常用中间件，轻松实现日志、认证、跨域、限流等功能。
<!-- - **统一错误处理**: 统一的错误处理机制，简化错误信息的捕获和处理。 -->

## 快速开始

```sh
git clone https://github.com/sanjayheaven/go-gin-boilerplate.git
cd go-gin-boilerplate
go mod download
```

### 创建 githooks 软链接【推荐】

```sh
cd .git/hooks
ln -s ../../githooks/* .
```

> 提示：
>
> - 如何确认已经创建成功?
>
> 运行以下命令：
>
> ```sh
> ls -l . # 现在你的当前目录位置应该在 .git/hooks 目录下
> ```
>
> 如果成功，你将会看到输出包含以下内容：
>
> ```sh
> commit-msg -> ../../githooks/commit-msg
> pre-commit -> ../../githooks/pre-commit
> ```

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
go run main.go server
```

## 打包

```sh
make build
```

## 部署

### docker-compose

使用 docker-compose 部署应用程序。  
确保 在服务器中已经安装 docker ，并且熟悉 docker compose 的使用

- 复制 `deployments/docker-compose.yml` 配置文件到 项目根目录 中
- 执行以下命令，启动应用程序

```sh
docker compose up -d
```

## 支持 🫶

- Star 🌟 项目
- 欢迎提交 [issue](https://github.com/sanjayheaven/go-gin-boilerplate/issues)。感谢您的支持
- 帮助在社交媒体上宣传并向朋友推荐它

  [![Twitter](https://img.shields.io/twitter/url?label=Twitter&logo=twitter&style=flat&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)
  [![Facebook](https://img.shields.io/twitter/url?label=Facebook&logo=facebook&style=flat&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)](https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)
  [![WhatsApp](https://img.shields.io/twitter/url?label=WhatsApp&logo=whatsapp&style=flat&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)](https://api.whatsapp.com/send?text=Wow:%20https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)
  [![Telegram](https://img.shields.io/twitter/url?label=Telegram&logo=telegram&style=flat&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)](https://t.me/share/url?url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)

- 你也可以在 [Ko-Fi](https://ko-fi.com/dorvan) 或者 [Buy Me A Coffee](https://www.buymeacoffee.com/dorvan) 上赞助一杯咖啡

  <a href='https://ko-fi.com/J3J1T95FG' target='_blank'>
  <img width="145" height="40" src='https://storage.ko-fi.com/cdn/kofi2.png?v=3' border='0' alt='Buy Me a Coffee at ko-fi.com' />
  </a>

  <a href="https://www.buymeacoffee.com/dorvan" target="_blank">
  <img width="145" height="40" src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" />
  </a>
