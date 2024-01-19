package logger

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var (
	ZapLogger *zap.Logger
	ZapSugar  *zap.SugaredLogger

	LogrusLogger *logrus.Logger

	// export Logger
	Logger *logrus.Logger
)

func init() {
	Logger = InitLogrusLogger()

	LogrusLogger = Logger
}
