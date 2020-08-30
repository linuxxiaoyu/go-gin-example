package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	log *logrus.Logger
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	log = new(logrus.Logger)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006/01/02 15:04:05",
	})
	log.SetOutput(F)
	log.SetLevel(logrus.InfoLevel)
	// logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	log.Debug(v)
}

func Info(v ...interface{}) {
	log.Info(v)
}

func Warn(v ...interface{}) {
	log.Warn(v)
}

func Error(v ...interface{}) {
	log.Error(v)
}

func Fatal(v ...interface{}) {
	log.Fatal(v)
}
