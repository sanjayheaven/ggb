# Go Gin Boilerplate

<div>

<img src="./assets/golang.png" align="right"/>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sanjayheaven/go-gin-boilerplate)
![CI](https://github.com/sanjayheaven/go-gin-boilerplate/workflows/Go/badge.svg)
![License](https://img.shields.io/github/license/sanjayheaven/go-gin-boilerplate)

[Swagger Api Document](https://go-gin-boilerplate.gganbu.services/swagger/index.html)

</div>

## Introduction

A development boilerplate based on the Gin framework, aimed at helping developers quickly build and develop web applications.

- [English](./README_en.md)
- [中文介绍](./README.md)

## Keywords

[Go](https://github.com/golang/go)
[Gin](https://github.com/gin-gonic/gin)
[Viper](https://github.com/spf13/viper)
[Gorm](https://github.com/go-gorm/gorm)
[Gin-Swagger](https://github.com/swaggo/gin-swagger)
[Air](https://github.com/cosmtrek/air)
[Logrus](https://github.com/sirupsen/logrus)
[Lumberjack](https://github.com/natefinch/lumberjack)
[Jwt](https://github.com/golang-jwt/jwt)

## Features

- **Rapid Development**: Accelerate project development and iteration using the Gin framework and related tools.
- **Simple and Easy to Use**: Follow the [project-layout](https://github.com/golang-standards/project-layout/tree/master) specification, providing a clear and simple code structure that even beginners can easily understand.
- **Integrated Logging System**: Integrated Logrus and Lumberjack for comprehensive log recording and management.
- **Hot Reload**: Utilize the Air tool for hot reloading, improving development efficiency.
- **Database Support**: Integrated Gorm, supporting mainstream databases such as MySQL, PostgreSQL, etc.
- **Flexible Middleware**: Incorporate common middleware for easy implementation of features such as logging, authentication, cross-origin resource sharing (CORS), rate limiting, etc.

## Quick Start

```sh
git clone https://github.com/sanjayheaven/go-gin-boilerplate.git
cd go-gin-boilerplate
go mod download
```

### Configuration Setup

- Navigate to the `configs` directory, copy the `config.example.yaml` file, and rename it to `config.yaml`.

```sh
cp configs/config.example.yaml configs/config.yaml
```

- Modify the configuration items in the `config.yaml` file.

```sh
vi configs/config.yaml
```

### Run the Application

- Run the project using [air](https://github.com/cosmtrek/air) **【Recommended】**

```sh
air
```

- Run the project using [go run](https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program)

```sh
go run cmd/main.go
```

## Deployment

### docker-compose

Deploy the application using **_docker-compose_**.
Ensure that Docker is installed on the server, and you are familiar with the usage of docker compose.

- Copy the [docker-compose.yml](./deployments/docker-compose.yml) configuration file to the project root directory.
- Execute the following command to start the application

```sh
docker compose up -d
```

## Support 🫶

- Star 🌟 the project
- Welcome to submit [issue](https://github.com/sanjayheaven/go-gin-boilerplate/issues). Thank you for your support
- Help promote it on social media and recommend it to friends

  [![Twitter](https://img.shields.io/twitter/url?label=Twitter&logo=twitter&style=flat&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)
  [![Facebook](https://img.shields.io/twitter/url?label=Facebook&logo=facebook&style=flat&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)](https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)
  [![WhatsApp](https://img.shields.io/twitter/url?label=WhatsApp&logo=whatsapp&style=flat&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)](https://api.whatsapp.com/send?text=Wow:%20https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)
  [![Telegram](https://img.shields.io/twitter/url?label=Telegram&logo=telegram&style=flat&url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)](https://t.me/share/url?url=https%3A%2F%2Fgithub.com%2Fsanjayheaven%2Fgo-gin-boilerplate)

- You can also sponsor a cup of coffee on [Ko-Fi](https://ko-fi.com/dorvan) or [Buy Me A Coffee](https://www.buymeacoffee.com/dorvan)

  <a href='https://ko-fi.com/J3J1T95FG' target='_blank'><img style='border:0px;height: 40px !important;width: 145px !important;' src='https://storage.ko-fi.com/cdn/kofi2.png?v=3' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>
  <a href="https://www.buymeacoffee.com/dorvan" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 40px !important;width: 145px !important;" ></a>
