---
sidebar_position: 0.9
---

# 配置

配置功能是软件开发当中常用的一个功能，在多环境中，例如开发环境、测试环境、生产环境等等，我们需要不同的配置，这时候就需要配置文件来管理。

同时，为了保持配置文件的安全性，我们通常会将配置环境文件保存在各个独立的环境中，通过加载文件，来获取对应的配置内容。

## 配置目录

遵循 [project-layout](https://github.com/golang-standards/project-layout) 规范，我们将 `配置文件模板或默认配置` 保存在 `configs` 目录下。

在 `configs` 目录下，我们可以看到以下文件：

```sh
configs
├── config.example.yaml
├── config.go
├── config.yaml
├── jwt.go
├── mongo.go
├── mysql.go
├── redis.go
└── server.go
```

- `config.example.yaml` 为配置文件示例，供参考
- `config.yaml` 为运行时项目加载的配置文件
- `config.go` 文件，这个文件是用来解析配置文件之后，保存配置内容的，也是 `config` 包的入口

而剩下的，则是各个工具或者组件的配置文件，例如 `jwt.go` 是 `JWT` 的配置文件，`mongo.go` 是 `MongoDB` 的配置文件，`mysql.go` 是 `MySQL` 的配置文件，`redis.go` 是 `Redis` 的配置文件，`server.go` 是`项目服务`的配置文件。

## 解析配置文件

我们在这里，通过 [Viper](https://github.com/spf13/viper) 来解析配置文件。

Viper 是 Go 应用程序的完整配置解决方案，可以处理所有类型的配置需求和格式，功能十分强大，它支持：

- 设置默认值
- 从 JSON、TOML、YAML、HCL 和 Java 属性文件中读取配置数据
- 实时监控和重新加载配置文件
- 从环境变量中读取
- 从远程配置系统（etcd 或 Consul）中读取
- 从命令行参数中读取
- 从缓冲区中读取
- 从密钥/值存储中读取
- 调用远程配置系统（etcd 或 Consul）以存储配置数据

## Viper 的应用

在 `configs/config.go` 文件中，我们通过 `init()` 函数来初始化配置文件，并将配置内容赋于变量 **EnvConfig**，代码如下：

```go
package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql  Mysql
	Redis  Redis
	Server Server
	Jwt    Jwt
	Mongo  Mongo
}

var EnvConfig *Config

func LoadConfig() *Config {

	path, err := os.Getwd() // get curent path
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/configs") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}

func init() {
	EnvConfig = LoadConfig()
	fmt.Printf("👻 EnvConfig: %+v\n", EnvConfig)
}

```

## 配置文件示例

目前启动 Go-Gin-Boilerplate 项目，至少需要 **MySQL** 和 **Redis** 两个服务的配置。

在 `configs/config.example.yaml` 文件中，我们可以看到以下内容：

```yaml
# HERE IS JUST THE CONFIGURATION EXAMPLE FOR THE APP
# DO NOT USING THIS FILE AS YOUR CONFIGURATION DIRECTLY
# PLEASE COPY THIS FILE TO configs/config.yaml AND MODIFY IT

# server
server:
  port: ":YOUR_SERVER_PORT"

# jwt
jwt:
  secret: YOUR_JWT_SECRET

# mysql
mysql:
  host: YOUR_MYSQL_HOST
  port: YOUR_MYSQL_PORT
  user: YOUR_MYSQL_USER
  password: YOUR_MYSQL_PASSWORD
  database: YOUR_DATABASE_NAME

# redis
redis:
  host: YOUR_REDIS_HOST
  port: YOUR_REDIS_PORT
  password: YOUR_REDIS_PASSWORD
  db: YOUR_REDIS_DB
```

在克隆项目之后，我们也可以非常方便的根据 `config.example.yaml` 文件，创建 `config.yaml` 文件，并修改其中的配置内容。

```sh
cp configs/config.example.yaml configs/config.yaml
```

修改 `config.yaml` 文件中的配置内容，将变量值替换为你自己的配置内容。

```sh
vi configs/config.yaml
```

## 使用配置包

配置包的入口是 `configs/config.go` 文件，因为我们已经在 `init()` 函数中，初始化了配置文件，所以我们可以在项目其他地方，通过 `configs.EnvConfig` 来获取配置内容。

例如在 项目 启动程序中，我们可以直接调用 `configs.EnvConfig` 来获取配置内容，代码如下：

```go
package main

import (
    "fmt"

    "github.com/sanjayheaven/ggb/configs"
)


func main() {
    fmt.Printf("👻 EnvConfig: %+v\n", configs.EnvConfig)
}
```
