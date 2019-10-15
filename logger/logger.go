package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger new logger to file
func NewLogger(logfile string) *logrus.Logger {
	logger := logrus.New()
	_, err := os.Stat(logfile)
	if err != nil || !os.IsExist(err) {
		os.Create(logfile)
	}
	fs, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	logger.Out = fs
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	return logger
}
