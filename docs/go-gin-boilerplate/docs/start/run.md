---
sidebar_position: 3
---

# 运行项目

在本地运行项目

## 使用 air 运行项目【开发环境推荐】

> [Air](https://github.com/cosmtrek/air) 是为 Go 应用开发设计的另外一个热重载的命令行工具。只需在你的项目根目录下输入 air，然后把它放在一边，专注于你的代码即可

在项目根目录中，打开终端，执行命令

```sh
air
```

air 配置文件 `.air.toml` 如下

```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
args_bin = []
bin = "./tmp/ggb server"
cmd = "go build -o ./tmp/ggb ./main.go"
delay = 0
exclude_dir = ["assets", "tmp", "vendor", "testdata"]
exclude_file = []
exclude_regex = ["_test.go"]
exclude_unchanged = false
follow_symlink = false
full_bin = ""
include_dir = []
include_ext = ["go", "tpl", "tmpl", "html"]
include_file = []
kill_delay = "0s"
log = "build-errors.log"
poll = false
poll_interval = 0
rerun = false
rerun_delay = 500
send_interrupt = false
stop_on_error = false

[color]
app = ""
build = "yellow"
main = "magenta"
runner = "green"
watcher = "cyan"

[log]
main_only = false
time = false

[misc]
clean_on_exit = false

[screen]
clear_on_rebuild = false
keep_scroll = true

```

## 使用 go run 运行项目

在项目根目录中，打开终端，执行命令

```sh
go run main.go server
```

## 打包后执行 【生产环境推荐】

### 打包程序

在终端执行命令

```sh
make build
```

或者

```sh
go build -o build/ggb main.go
```

### 运行程序

打包后的程序路径为 build/ggb, 在终端执行命令

```sh
./build/ggb server
```

## 查看更多命令

```sh
go run main.go --help
```
