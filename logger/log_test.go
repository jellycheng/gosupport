package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInfo(t *testing.T) {

	//级别
	logrus.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logrus.SetFormatter(&SimpleTextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		LogFormat: "%level% %time% %msg% %app_name%\n",
	})


	logrus.Debug("debug信息")
	logrus.Info("info信息")

	logrus.WithField("app_name", "user-service").Error("Log message")

	logrus.WithField("app_name", 300.123).Error("Log message")

}


