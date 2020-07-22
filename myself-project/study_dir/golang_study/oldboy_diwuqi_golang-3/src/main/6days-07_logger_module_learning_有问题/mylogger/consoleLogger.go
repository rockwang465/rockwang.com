package mylogger

import (
	"fmt"
	"time"
)

// 1 创建class
type ConsoleLogger struct {
	// 4.2 保存用户传入的日志等级，并已经转为int16格式了
	logLevel int16
}

// 2 创建构造函数
func NewConsoleLog(logLevelStr string) ConsoleLogger {
	// 4 读取传入的日志等级，转换成数字
	logLevelInt, err := logLevelToInt(logLevelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		// 4.1 保存用户传入的日志等级，并已经转为int16格式了
		logLevelInt, // Rock注意1: 这里总忘记加逗号,
	}
}

// 5 判断用户输入的日志等级，和当前日志函数等级，确认是否进行日志打印
func compareLevel(logLevel int16, logLevelStr, formatMsg string, a ...interface{}) {
	// 6.2 支持格式化打印日志("id:%s name:%s", 2, rock)
	msg := fmt.Sprintf(formatMsg, a...)

	nowTime := time.Now()
	layout := "2006-01-02 15:04:05"
	nowTimeF := nowTime.Format(layout)
	//time.Sleep(time.Second * 1)
	curLevel, _ := logLevelToInt(logLevelStr)
	// 6.2 获取ERROR级别的int值，进行判断不同等级，打印不同格式的日志
	errorLogLevel, _ := logLevelToInt("ERROR")
	// 6.1 获取报错的函数名、文件名、行号信息
	funcName, fileName, line := callUser(3)
	if logLevel <= curLevel {
		// 6.3 如果是 >= error级别日志，则打印报错的文件名、函数名、行号； 反之不打印
		if curLevel >= errorLogLevel {
			//fmt.Println(errorLogLevel)
			fmt.Printf("[%s %s %s:%s:%d] %s\n", nowTimeF, logLevelStr, fileName, funcName, line, msg)
		} else {
			fmt.Printf("[%s %s] %s\n", nowTimeF, logLevelStr, msg)
		}
	}
}

// 3 生成对应调用的方法
// 6.1 支持格式化打印日志("id:%s name:%s", 2, rock), 这里则传参为(formatMsg string, a ...interface{})
func (c ConsoleLogger) Debug(formatMsg string, a ...interface{}) {
	//curLevel,_ := logLevelToInt("Debug")
	//if c.logLevel <= curLevel{
	//	fmt.Printf("[Debug ] %s\n", msg)
	//}
	//  5.1 传参、执行等级判断，并打印日志
	compareLevel(c.logLevel, "DEBUG", formatMsg, a...)
}
func (c ConsoleLogger) Info(formatMsg string, a ...interface{}) {
	//fmt.Printf("[Info ] %s\n", msg)
	compareLevel(c.logLevel, "INFO", formatMsg, a...)
}
func (c ConsoleLogger) Error(formatMsg string, a ...interface{}) {
	//fmt.Printf("[Error ] %s\n", msg)
	compareLevel(c.logLevel, "ERROR", formatMsg, a...)
}
