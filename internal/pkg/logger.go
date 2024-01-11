package Logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger

func Init() *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.SetReportCaller(true)

	logFile := &lumberjack.Logger{
		Filename: "logs/app.log",
		MaxSize:  10,
		Compress: true,
	}

	// multi writer, both file and stdout
	writers := []io.Writer{logFile, os.Stdout}
	logger.SetOutput(io.MultiWriter(writers...))

	// export Logger
	Logger = logger

	return logger
}
