package main

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	//logrus.SetLevel(logrus.TraceLevel)  // 设置日志等级
	logrus.SetLevel(logrus.InfoLevel) // 设置日志等级

	logrus.SetFormatter(&nested.Formatter{ // 格式化日志
		// HideKeys:        true,
		TimestampFormat: time.RFC3339,
		FieldsOrder:     []string{"name", "age"},
	})

	//logrus.WithFields(logrus.Fields{
	//	"name": "dj",
	//	"age":  18,
	//}).Info("info msg")

	logrus.Trace("trace msg") // 普通日志打印
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	//logrus.Fatal("fatal msg")
	//logrus.Panic("panic msg")
}
