---
sidebar_position: 3
---

# 运行

在本地运行项目

## 使用 air 运行项目【推荐】

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

```sh
go run main.go server
```

## 查看更多命令

```sh
go run main.go --help
```
