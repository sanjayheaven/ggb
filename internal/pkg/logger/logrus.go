package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

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
