---
sidebar_position: 4
---

# 正式部署

多种方式部署你的项目

## docker-compose

### 要求

- 已有 docker 环境
- 已有 mysql、redis 等服务运行

1. 修改 `configs/config.yaml` 文件中的配置项，确保 mysql、redis 等服务的配置正确，并复制到项目根目录 中

### 配置

1. 修改 `configs/config.yaml` 文件中的配置项，确保 mysql、redis 等服务的配置正确，并复制到项目根目录 中
2. 修改 `deployments/docker-compose.yml` 文件中的配置项，并复制到项目根目录 中

### 运行项目

在项目根目录下，执行以下命令，启动项目

```sh
docker compose up -d
```

### docker-compose 示例

```yml
version: "3"
services:
  go-gin-boilerplate:
    image: golang:latest
    container_name: go-gin-boilerplate
    working_dir: /server
    volumes:
      - ./:/server # mount current directory to /server, make sure your docker-compose.yml is in the root of your project
    restart: always
    command: sh -c "make build && ./build/ggb server"
```

## 二进制文件运行

### 配置

1. 修改 `configs/config.yaml` 文件中的配置项，确保 mysql、redis 等服务的配置正确，并复制到项目根目录 中

### 打包

```sh
make build
```

### 执行二进制文件

```sh
./build/ggb server
```
