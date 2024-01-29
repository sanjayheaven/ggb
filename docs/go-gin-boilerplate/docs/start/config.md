---
sidebar_position: 2
---

# 环境配置

创建配置文件，用于配置项目的基本信息。

## 创建本地配置文件

复制 `configs/config.example.yaml` 到 `configs/config.yaml`，并根据需要进行修改。

```sh
cp configs/config.example.yaml configs/config.yaml
```

参考：`config.example.yaml` 文件内容如下

```yml
# HERE IS JUST THE CONFIGURATION EXAMPLE FOR THE APP
# DO NOT USING THIS FILE AS YOUR CONFIGURATION DIRECTLY
# PLEASE COPY THIS FILE TO configs/config.yaml AND MODIFY IT

# server
server:
  port: ":8080"
```
