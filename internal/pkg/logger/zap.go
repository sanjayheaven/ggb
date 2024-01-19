package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

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
