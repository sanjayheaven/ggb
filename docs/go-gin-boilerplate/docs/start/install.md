---
sidebar_position: 1
---

# 项目初始化

如何在本地安装项目，并且下载依赖，完成项目的初始化准备工作。

## 要求

- [Go](https://golang.org/) >= 1.20

## 安装

### 克隆项目

```sh
git clone https://github.com/sanjayheaven/ggb.git
```

### 安装依赖

进入工作目录

```sh
cd go-gin-boilerplate
```

执行以下命令，安装依赖

```sh
go mod download
```

## 创建 githooks 软链接【推荐】

通过创建软链接，可以在每次提交代码时，自动运行 `pre-commit` 和 `commit-msg` 钩子，从而实现自动化测试和规范化提交信息。

进入 `.git/hooks` 目录

```sh
cd .git/hooks
```

建立软链接

```sh
ln -s ../../githooks/* .
```

### 提示

> - 如何确认已经创建成功?
>
> 运行以下命令：
>
> 现在你的当前目录位置应该在 .git/hooks 目录下
>
> ```sh
> ls -l .
> ```
>
> 如果成功，你将会看到类似输出会包含以下内容：
>
> ```sh
> commit-msg -> ../../githooks/commit-msg
> pre-commit -> ../../githooks/pre-commit
> ```
