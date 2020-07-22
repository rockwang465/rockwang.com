package mylogger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

// 15 interface连接ConsoleLogger 和 FileLogger 两个结构体,实现多态功能,成为整个日志模块的总入口*****
type Logger interface {
	Debug(formatMsg string, a ...interface{})
	Info(formatMsg string, a ...interface{})
	Error(formatMsg string, a ...interface{})
	Fatal(formatMsg string, a ...interface{})
}

// 注: 此go文件用于向文件中输出报错日志

// 13.1 定义输入到文件的所需结构体
type FileLogger struct {
	LogLevel      logLevelInt // 用户选择的日志输出等级
	fileName      string      // 输出的日志文件名
	filePath      string      // 输出的日志文件路径
	maxLogSize    int64       // 定义日志最大的大小
	logFileObj    *os.File    // 打开日志文件的文件对象(普通日志)
	errLogFileObj *os.File    // 打开日志文件的文件对象(错误日志)
}

// 13.2 定义构造函数
func NewFileLog(lv, fileName, filePath string, maxLogSize int64) (fl *FileLogger) {
	// 日志等级字符串转为数字
	lvInt, err := LogLevelToInt(lv)
	if err != nil {
		panic(err)
	}

	fl = &FileLogger{
		LogLevel:   lvInt,
		fileName:   fileName,
		filePath:   filePath,
		maxLogSize: maxLogSize,
	}

	// 13.3.1 open 日常日志文件
	// 13.3.2 open 错误日志文件
	err = fl.initFile()
	if err != nil {
		fmt.Println("Error : initFile was failed : ", err)
		return
	}

	return fl
}

// 13.3.3 打开用户传入的文件
func (fl *FileLogger) initFile() error {
	//拼接路径(优点: 可以用于windows和linux，因为拼接的/\ 在linux下和windows下不同，有识别功能)
	fullFilePath := path.Join(fl.filePath, fl.fileName)
	// 日常日志文件打开，生成文件对象，用于操作
	logFileObj, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error : open logFileObj file was failed ")
		return err
	}

	// 错误日志文件打开，生成文件对象，用于操作
	errLogFileObj, err := os.OpenFile(fullFilePath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error : open errLogFileObj file was failed ")
		return err
	}

	fl.logFileObj = logFileObj
	fl.errLogFileObj = errLogFileObj
	return nil
}

// 13.3.4关闭open的文件 -- 这个功能用不上，因为需要切割日志的时候才会关闭当前的文件对象
//func (fl *FileLogger) Close() {
//	fl.logFileObj.Close()
//	fl.errLogFileObj.Close()
//}

// 14.1 超过则切割日志文件
func (fl *FileLogger) checkLogSize(file *os.File) bool {
	FileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error : check log size use file.Stat err = ", err)
		return false
	}
	fileSize := FileInfo.Size()
	if fileSize >= fl.maxLogSize { // 如果当前日志文件已经大于等于用户定义的大小，则返回true(true则需要切割日志)
		return true
	} else {
		return false
	}
	//return FileInfo.Size() >= fl.maxLogSize   // 老师简写，比上面简单太多了
}

// 14.2 检查当前日志文件大小是否超过预定义大小
func (fl *FileLogger) cutLog(file *os.File) (fileObj *os.File, err error) {
	// A.获取当前传入进来的文件对象的信息，拿到文件名、文件路径
	FileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat failed , err = ", err)
		return nil, err // 老师这里返回的内容
	}
	fileName := FileInfo.Name() // 当前传入进来的文件对象的文件名(不用管是普通日志文件 还是 err日志文件 了)
	fullFilePath := path.Join(fl.filePath, fileName)
	fmt.Printf("要切割的日志文件全路径: %s\n", fullFilePath)

	// B.定义需要备份的文件名(.bak+时间戳)
	_, nowTime := getNowTime()
	newBakFilePath := fmt.Sprintf("%s.bak-%s", fullFilePath, nowTime) // 定义要备份的文件名(.bak+时间戳)
	fmt.Printf("要重命名的日志文件全路径: %s\n", newBakFilePath)

	// C.关闭当前文件对象
	file.Close() // 注意: 关闭了文件对象，就无法获取对应的文件信息了，所以上面先获取，这里才能关闭。

	// D.重命名文件
	err = os.Rename(fullFilePath, newBakFilePath) // 重命名，来备份原文件
	if err != nil {
		fmt.Println("os.Rename failed, err = ", err)
		return nil, err
	}

	// D.open新的日志文件
	fileObj, err = os.OpenFile(fullFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("os open new log  file failed, err = ", err)
		return nil, err
	}
	return fileObj, nil
}

// +----------------------------------------------------------------------------------------------------+
// + 注意: 以下部分,全部copy自consoleLogger.go,因为功能相同,所以copy过来,替换下结构体名称及写入到文件功能就行   +
// +----------------------------------------------------------------------------------------------------+

func (fl *FileLogger) CompareLogLevel(FuncLogLevel logLevelInt) bool {
	if fl.LogLevel <= FuncLogLevel { // 7.2 当函数级别大于等于用户输入级别，为真，则返回true，打印该级别日志
		return true
	} else {
		return false // 当函数级别小于用户输入级别，为假，则返回false，不打印该级别日志
	}
}

// 9 提示用户报错的文件名、函数、报错行号
func (fl *FileLogger) callUser(skip int) (fileName, funcName string, line int) { // Rock注意4: 这里必须写好每个返回值的类型。
	// skip :int类型，是用于当前函数fl.callUser之前被哪个函数调用，计算出层数。 例如fl.callUser被f1()调用，f1()被f2，得出被调用2次，那么这个skip的值应该为2.
	// 这样，当输出给用户报错信息时，就能定位到当前报错，被第一个函数调用的位置了。

	// pc :函数的信息保存，主要用于取当前的函数名
	// file :当前logger.go文件的绝对路径文件名
	// line :显示当前函数被skip次(被调用的N次)前函数的行号。也就是skip的数值，对应前面第几层调用当前函数的行号。
	// ok :如果能够取到(runtime.Caller能正常执行)，则为true
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("Error : runtime.Caller failed")
	}
	fileName = path.Base(file)
	//funcName = runtime.FuncForPC(pc).Name() // main.main
	lenFuncName := len(strings.Split(runtime.FuncForPC(pc).Name(), "."))       // 先取长度
	funcName = strings.Split(runtime.FuncForPC(pc).Name(), ".")[lenFuncName-1] // 从结果main.main中，取后面一个函数名就好，不然太多有点难看
	//fmt.Println("当前文件的文件名: ", file)   // E:/3.常用代码目录/go_study/老男孩第五期golang视频课程-3/src/main/6days-07_logger_module/server.go
	//fmt.Println("当前调用的行号: ", line)    // 19 (server.go，第19行，调用Fatal函数，Fatal函数调用了fl.callUser函数。 skip值为2)
	//fmt.Println("当前文件名: ", fileName)  // server.go
	//fmt.Println("当前执行的函数名", funcName) // main.main
	return
}

// 11 将日志函数中一大堆的重复内容，放到这里，后面直接调用
//func fl.repeatInfo(lvInt logLevelInt, msg string) {
// 12 支持msg为格式化的内容传参,且支持格式化的内容中有多个格式化的值("id:%d, name:%s", id ,name 这种方式)
func (fl *FileLogger) repeatInfo(lvInt logLevelInt, formatMsg string, a ...interface{}) {
	// 12.1 字符串格式化后，存入msg中
	msg := fmt.Sprintf(formatMsg, a...) // Rock注意5: 格式化字符串的方法

	nowTime, _ := getNowTime()
	// 9.1 报错日志，正常都需要报错的文件名、函数名、行号等信息，所以这里就是提供这些。
	fileName, funcName, line := fl.callUser(3)
	// 10.1 将DEBUG(数字类型)，转成字符串类型，返回过来，放入下面fmt.Printf中使用
	strLogLevel := LogLevelToStr(lvInt)

	// 14.3 检查当前日志文件大小是否超过预定义大小，
	resBool := fl.checkLogSize(fl.logFileObj)
	if resBool { // 如果返回值为true，则需要切割日志
		// 14.4 超过则切割日志文件
		fileObj, err := fl.cutLog(fl.logFileObj)
		if err != nil {
			fmt.Println("Error : open file failed , err = ", err)
			return
		}
		fl.logFileObj = fileObj // 将新打开的文件的文件对象赋值给结构体中的普通文件对象
	}
	fmt.Fprintf(fl.logFileObj, "[%s %s] [%s:%s:%v] : %s\n", nowTime, strLogLevel, fileName, funcName, line, msg)

	if lvInt >= ERROR {
		// 14.3 检查当前err日志文件大小是否超过预定义大小
		resBool := fl.checkLogSize(fl.errLogFileObj)
		if resBool { // 如果返回值为true，则需要切割日志
			// 14.4 超过则切割日志文件
			fileObj, err := fl.cutLog(fl.errLogFileObj)
			if err != nil {
				fmt.Println("Error : open error file failed, err = ", err)
				return
			}
			fl.errLogFileObj = fileObj // 将新打开的文件的文件对象赋值给结构体中的err日志文件对象
		}
		fmt.Fprintf(fl.errLogFileObj, "[%s %s] [%s:%s:%v] : %s\n", nowTime, strLogLevel, fileName, funcName, line, msg)
	}
}

// 3.Debug 方法，日志打印
//func (fl *FileLogger) Debug(msg string) {
// 12.2 更换格式化传参
func (fl *FileLogger) Debug(formatMsg string, a ...interface{}) {
	// 7.判断用户传入日志级别和当前函数的级别大小，这个有点绕
	// 可以这么理解: 如果用户输入的Error级别的，而当前函数为Debug级别，如果判断为false，则不会打印;否则弄反了，Error级别就会打印Debug级别日志了。
	// 所以应该用 log.Level <= DEBUG
	//if log.LogLevel <= DEBUG { // 这里可以做成一个函数来判断更省事
	if resBool := fl.CompareLogLevel(DEBUG); resBool { // 7.3 这里判断返回值，true则打印，false则不打印
		// 11.1 注释掉一下重复的内容，将重复的内容让repeatInfo函数去做
		//nowTime := getNowTime()
		//// 9.1 报错日志，正常都需要报错的文件名、函数名、行号等信息，所以这里就是提供这些。
		//fileName, funcName, line := fl.callUser(2)
		//// 10.1 将DEBUG(数字类型)，转成字符串类型，返回过来，放入下面fmt.Printf中使用
		//strLogLevel := LogLevelToStr(DEBUG)
		//fmt.Printf("[%s %s] [%s:%s:%v] : %s\n", nowTime, strLogLevel, fileName, funcName, line, msg)

		// 11.2 调用repeatInfo函数，将重复的内容全部交给repeatInfo去做
		//fl.repeatInfo(DEBUG, msg)
		// 12.2 支持msg为格式化的内容传参,且支持格式化的内容中有多个格式化的值("id:%d, name:%s", id ,name 这种方式)
		fl.repeatInfo(DEBUG, formatMsg, a...)
	}
}

// 3.Info 方法，日志打印
//func (fl *FileLogger) Info(msg string) {
// 12.2 更换格式化传参
func (fl *FileLogger) Info(formatMsg string, a ...interface{}) {
	if fl.CompareLogLevel(INFO) { // 7.4 这里简写判断为true/false
		// 11.2 调用repeatInfo函数，将重复的内容全部交给repeatInfo去做
		//fl.repeatInfo(INFO, msg)
		// 12.3 支持msg为格式化的内容传参,且支持格式化的内容中有多个格式化的值("id:%d, name:%s", id ,name 这种方式)
		fl.repeatInfo(INFO, formatMsg, a...)
	}
}

// 3.Error 方法，日志打印
//func (fl *FileLogger) Error(msg string) {
func (fl *FileLogger) Error(formatMsg string, a ...interface{}) {
	if fl.CompareLogLevel(ERROR) { // 7.4 这里简写判断为true/false
		// 11.2 调用repeatInfo函数，将重复的内容全部交给repeatInfo去做
		//fl.repeatInfo(ERROR, msg)
		fl.repeatInfo(ERROR, formatMsg, a...)
	}
}

// 3.Fatal 方法，日志打印
//func (fl *FileLogger) Fatal(msg string) {
func (fl *FileLogger) Fatal(formatMsg string, a ...interface{}) {
	if fl.CompareLogLevel(FATAL) { // 7.4 这里简写判断为true/false
		// 11.2 调用repeatInfo函数，将重复的内容全部交给repeatInfo去做
		//fl.repeatInfo(FATAL, msg)
		fl.repeatInfo(FATAL, formatMsg, a...)
	}
}
