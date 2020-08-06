package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 7 结构体
type FileLogger struct {
	logLevel      int16
	logName       string
	errLogName    string
	filePath      string
	maxFileSize   int64
	logFileObj    *os.File
	errLogFileObj *os.File
}

// 8 构造函数
func NewFileLog(logLevelStr, logName, errLogName, filePath string, maxFileSize int64) *FileLogger {
	logLevelInt, err := logLevelToInt(logLevelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		logLevel:      logLevelInt,
		logName:       logName,
		errLogName:    errLogName,
		filePath:      filePath,
		maxFileSize:   maxFileSize,
		logFileObj:    nil,
		errLogFileObj: nil,
	}

	fl.initFiles()

	return fl
}

// 8.1 初始化日志文件与err日志文件， 生成文件对象
func (fl *FileLogger) initFiles() {
	// 普通日志文件
	fullLogPath := path.Join(fl.filePath, fl.logName)
	fileObj, err := os.OpenFile(fullLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fl.logFileObj = fileObj

	// error日志文件
	fullErrLogPath := path.Join(fl.filePath, fl.errLogName)
	errFileObj, err := os.OpenFile(fullErrLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fl.errLogFileObj = errFileObj
}

// 9.1 检查日志的大小是否超过定义的日志大小
func (fl *FileLogger) checkSize(fileObj *os.File) bool {
	fmt.Println("1.进入了 检查日志大小 函数了")
	FileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Println("Error : check size Stat err = ", err)
		return false
	}

	curFileSize := FileInfo.Size()
	if curFileSize >= fl.maxFileSize {
		fmt.Println("2.确认超过定义的大小，需要进入日志切割函数进行切割了")
		return true // 如果超过定义的日志大小，则需要切割日志
	} else {
		return false
	}
}

// 9.3 日志切割
func (fl *FileLogger) cutLog(fileObj *os.File) (*os.File, error) {
	fmt.Println("3.进入了 切割日志函数了")
	FileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Println("Error : cut log Stat err = ", err)
		return nil, err
	}

	fileName := FileInfo.Name()
	fullLogPath := path.Join(fl.filePath, fileName)
	fmt.Printf("3.1要切割的文件全路径为: %s\n", fullLogPath)

	NT := time.Now()
	curTime := NT.Format("20060504150201")
	bakLogPath := fmt.Sprintf("%s.bak-%s", fullLogPath, curTime)
	fmt.Printf("3.2要重命名的文件全路径为: %s\n", bakLogPath)

	// 关闭当前打开的文件对象
	fileObj.Close()

	// 重命名
	err = os.Rename(fullLogPath, bakLogPath)
	if err != nil {
		fmt.Println("Error : os rename err")
		panic(err)
	}

	// 打开新的文件出来
	fileObj, err = os.OpenFile(fullLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error : os open new log file err = ", err)
		return nil, err
	}
	fmt.Println("4.已切割完成日志文件，并重新打开源文件名，返回文件对象")
	return fileObj, nil
}

// 5 判断用户输入的日志等级，和当前日志函数等级，确认是否进行日志打印
func (fl *FileLogger) compareLevel(logLevel int16, logLevelStr, formatMsg string, a ...interface{}) {
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
	if logLevel <= curLevel { // 判断是否打印日志
		//9.2 判断日志大小，超过则备份日志(重命名原文件，然后创建新的日志文件为之前的文件名)
		if fl.checkSize(fl.logFileObj) {
			// 9.4 切割日志
			fileObj, err := fl.cutLog(fl.logFileObj) // 如果达到定义的日志大小，则需要切割日志
			if err != nil {
				fmt.Println("Error : cut formal log err = ", err)
				return
			}
			fl.logFileObj = fileObj
		}
		//fmt.Printf("[%s %s] %s\n", nowTimeF, logLevelStr, msg)
		fmt.Fprintf(fl.logFileObj, "[%s %s] %s\n", nowTimeF, logLevelStr, msg)

		// 6.3 如果是 >= error级别日志，则打印报错的文件名、函数名、行号； 反之不打印
		if curLevel >= errorLogLevel {
			//9.2 判断日志大小，超过则备份日志(重命名原文件，然后创建新的日志文件为之前的文件名)
			if fl.checkSize(fl.errLogFileObj) {
				// 9.4 切割日志
				fileObj, err := fl.cutLog(fl.errLogFileObj) // 如果达到定义的日志大小，则需要切割日志
				if err != nil {
					fmt.Println("Error : cut error log err = ", err)
					return
				}
				fl.errLogFileObj = fileObj
			}

			//fmt.Println(errorLogLevel)
			//fmt.Printf("[%s %s %s:%s:%d] %s\n", nowTimeF, logLevelStr, fileName, funcName, line, msg)
			fmt.Fprintf(fl.errLogFileObj, "[%s %s %s:%s:%d] %s\n", nowTimeF, logLevelStr, fileName, funcName, line, msg)
		}
	}
}

// 3 生成对应调用的方法
// 6.1 支持格式化打印日志("id:%s name:%s", 2, rock), 这里则传参为(formatMsg string, a ...interface{})
func (fl *FileLogger) Debug(formatMsg string, a ...interface{}) {
	//  5.1 传参、执行等级判断，并打印日志
	fl.compareLevel(fl.logLevel, "DEBUG", formatMsg, a...)
}

func (fl *FileLogger) Info(formatMsg string, a ...interface{}) {
	fl.compareLevel(fl.logLevel, "INFO", formatMsg, a...)
}

func (fl *FileLogger) Error(formatMsg string, a ...interface{}) {
	fl.compareLevel(fl.logLevel, "ERROR", formatMsg, a...)
}
