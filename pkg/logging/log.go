package logging

import (
	"os"

	"github.com/linuxxiaoyu/go-gin-example/pkg/setting"

	"github.com/sirupsen/logrus"
)

type Level int

var (
	F *os.File

	log *logrus.Logger
)

func Setup() {
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err := openLogFile(fileName, filePath)
	if err != nil {
		logrus.Fatalln(err)
	}

	log = new(logrus.Logger)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: setting.AppSetting.TimeFormat,
	})
	log.SetOutput(F)
	log.SetLevel(logrus.InfoLevel)
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
