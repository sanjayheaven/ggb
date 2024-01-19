---
sidebar_position: 5
---

# 日志

日志是一个记录事件的过程，例如：记录应用程序的运行状态、记录用户的操作行为等。

## 介绍

Go Gin Boilerplate 集成了两种主流的日志库，分别是 [logrus](https://github.com/sirupsen/logrus) 和 [zap](https://github.com/uber-go/zap)。

## 需求

在开发过程中，我们需要记录应用程序的运行状态，例如：记录用户的操作行为、记录应用程序的运行状态等。

我们通常会有以下几点主要的需求：

- 日志输出到文件、控制台
- 日志分级，例如：`debug`、`info`、`warn`、`error`、`fatal`、`panic`
- 日志切割，按照时间或者大小切割

## 日志包

我们定义了一个日志包，在 `internal/pkg/logger` 目录下，用于封装日志库的使用。

其中 `logger.go` 作为日志包的入口，定义了日志的初始化方法 `init()`，用于初始化日志库。

默认选择 **logrus** 作为日志库，如果需要使用 zap 作为日志库，可以修改 `logger.go` 中的初始化方法。

```go
var (
	ZapLogger *zap.Logger
	ZapSugar  *zap.SugaredLogger

	LogrusLogger *logrus.Logger
)

func init() {
	LogrusLogger = InitLogrusLogger()

	// ZapLogger = InitZapLogger()
	// ZapSugar = ZapLogger.Sugar()
}
```

## logrus

logrus 是一个结构化的日志库，提供了丰富的日志输出格式，例如：JSON、Text、Logfmt 等。

在 `logger/logrus.go` 中，我们定义了一个初始化方法 `InitLogrusLogger()`，用于初始化 logrus 日志库，返回一个 `logrus.Logger` 实例。

```go
func InitLogrusLogger() *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logFile := &lumberjack.Logger{
		Filename: "logs/app.log",
		MaxSize:  10,
		Compress: true,
	}

	// multi writer, both file and stdout
	writers := []io.Writer{logFile, os.Stdout}
	logger.SetOutput(io.MultiWriter(writers...))

	return logger
}
```

## zap

zap 是一个高性能的日志库，提供了丰富的日志输出格式，例如：JSON、Text、Logfmt 等。

在 `logger/zap.go` 中，我们定义了一个初始化方法 `InitZapLogger()`，用于初始化 zap 日志库，返回一个 `zap.Logger` 实例。

```go
func InitZapLogger() *zap.Logger {

	encoder := getEncoder()

	// First, define our level-handling logic.

	// level: debug,info,warning
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	// level: error, dpanic, panic, fatal
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	infoLevel_writerSyncer := getWriterSyncer("info")
	errorLevel_writerSyncer := getWriterSyncer("error")

	info_multiWriteSyncer := zapcore.NewMultiWriteSyncer(infoLevel_writerSyncer, os.Stdout)
	error_multiWriteSyncer := zapcore.NewMultiWriteSyncer(errorLevel_writerSyncer, os.Stdout)

	core := zapcore.NewCore(encoder, info_multiWriteSyncer, infoLevel)
	errorCore := zapcore.NewCore(encoder, error_multiWriteSyncer, errorLevel)

	coreArr := []zapcore.Core{core, errorCore}

	// export
	zapLogger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) // zap.AddCaller() will add line number and file name
	return zapLogger

}
```

其中 两个内部函数 `getEncoder()` 和 `getWriterSyncer()` 分别用于获取日志编码器和日志输出器。

```go

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer(level string) zapcore.WriteSyncer {
	lumberWriteSyncer := &lumberjack.Logger{
		Filename: fmt.Sprintf("logs/%s.log", level),
		MaxSize:  10, // megabytes
		Compress: true,
	}
	// file, _ := os.Create("logs/app.log")
	return zapcore.Lock(zapcore.AddSync(lumberWriteSyncer))

}
```

## lumberjack

lumberjack 是一个日志切割库，可以按照时间或者大小切割日志文件。

在 LogrusLogger 和 ZapLogger 中，我们都使用了 lumberjack 日志切割库。

```go
// logrus.go
logFile := &lumberjack.Logger{
    Filename: "logs/app.log",
    MaxSize:  10,
    Compress: true,
}
```

```go
// zap.go
lumberWriteSyncer := &lumberjack.Logger{
    Filename: fmt.Sprintf("logs/%s.log", level),
    MaxSize:  10, // megabytes
    Compress: true,
}
```
