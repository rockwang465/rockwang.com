package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 注: 此go文件用于向屏幕输出报错日志

// 1.Logger 结构体，用于构造函数做成对象，然后有 .xx  方法使用
type ConsoleLogger struct {
	// 5.5 增加了返回值，用于判断日志级别高低
	LogLevel logLevelInt
}

// 2.NewConsoleLog 构造函数，使用此构造函数返回Logger结构体
// 5.3 将日志级别放入到构造函数传参中来
func NewConsoleLog(logLevelStr string) ConsoleLogger { // Rock注意2: 这里Logger是结构体，因为你定义的时候是type，所以Logger本身就是类型，不要理解错了。
	//                                      Rock注意3: 不管任何返回值，都必须填入类型。你可以没有返回值名称，但必须填入返回值的类型
	//return ConsoleLogger{}  // 2.1 没参数用这个，现在有传参了，则需要用下面的return

	// 6.3 将函数放到这里，把参数传入后，直接拿到用户输入的字符串日志级别对应的数字日志级别，方便后面用 大于 等于 小于判断。
	logLevelIntVal, err := LogLevelToInt(logLevelStr)
	if err != nil {
		panic(err) // 6.4 如果异常，这里直接就报错吧，不返回了。
	}

	// 5.4 return 日志级别给结构体，让结构体拿到参数进行判断
	return ConsoleLogger{
		//LogLevel:	logLevelIntVal,  // 老师的写法，这里Rock看不懂为啥这样写
		logLevelIntVal, // Rock还是这样写吧，看的明白
	}
}

// 7.1 用函数方式进行判断当前日志级别和当前函数的日志级别，只是为了减少代码量(用于下面3中的Debug、Info、Error、Fatal函数中的判断)
// 老师这里，直接将比对作为Logger的一个子方法，这样确实很好 -- Rock注意3: 这里老师的用法很巧妙，就不需要传2个参数进来了
//func CompareLogLevel(FuncLogLevel logLevelInt, LogLevel logLevelInt) bool {  // Rock原本没想到作为Logger的子方法使用，所以传了2个参数进来
func (con ConsoleLogger) CompareLogLevel(FuncLogLevel logLevelInt) bool {
	if con.LogLevel <= FuncLogLevel { // 7.2 当函数级别大于等于用户输入级别，为真，则返回true，打印该级别日志
		return true
	} else {
		return false // 当函数级别小于用户输入级别，为假，则返回false，不打印该级别日志
	}
}

// 9 提示用户报错的文件名、函数、报错行号
func callUser(skip int) (fileName, funcName string, line int) { // Rock注意4: 这里必须写好每个返回值的类型。
	// skip :int类型，是用于当前函数callUser之前被哪个函数调用，计算出层数。 例如callUser被f1()调用，f1()被f2，得出被调用2次，那么这个skip的值应该为2.
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
	//fmt.Println("当前调用的行号: ", line)    // 19 (server.go，第19行，调用Fatal函数，Fatal函数调用了callUser函数。 skip值为2)
	//fmt.Println("当前文件名: ", fileName)  // server.go
	//fmt.Println("当前执行的函数名", funcName) // main.main
	return
}

// 11 将日志函数中一大堆的重复内容，放到这里，后面直接调用
//func repeatInfo(lvInt logLevelInt, msg string) {
// 12 支持msg为格式化的内容传参,且支持格式化的内容中有多个格式化的值("id:%d, name:%s", id ,name 这种方式)
func repeatInfo(lvInt logLevelInt, formatMsg string, a ...interface{}) {
	// 12.1 字符串格式化后，存入msg中
	msg := fmt.Sprintf(formatMsg, a...) // Rock注意5: 格式化字符串的方法

	nowTime, _ := getNowTime()
	// 9.1 报错日志，正常都需要报错的文件名、函数名、行号等信息，所以这里就是提供这些。
	fileName, funcName, line := callUser(3)
	// 10.1 将DEBUG(数字类型)，转成字符串类型，返回过来，放入下面fmt.Printf中使用
	strLogLevel := LogLevelToStr(lvInt)
	fmt.Printf("[%s %s] [%s:%s:%v] : %s\n", nowTime, strLogLevel, fileName, funcName, line, msg)
}

//// 4.获取当前时间
//func getNowTime() string {
//	NT := time.Now()
//	// 注意 "2006-01-02 15:04:05" 这里是固定的，不可以随便写
//	return NT.Format("2006-01-02 15:04:05")
//}

// 3.Debug 方法，日志打印
//func (con ConsoleLogger) Debug(msg string) {
// 12.2 更换格式化传参
func (con ConsoleLogger) Debug(formatMsg string, a ...interface{}) {
	// 7.判断用户传入日志级别和当前函数的级别大小，这个有点绕
	// 可以这么理解: 如果用户输入的Error级别的，而当前函数为Debug级别，如果判断为false，则不会打印;否则弄反了，Error级别就会打印Debug级别日志了。
	// 所以应该用 log.Level <= DEBUG
	//if log.LogLevel <= DEBUG { // 这里可以做成一个函数来判断更省事
	if resBool := con.CompareLogLevel(DEBUG); resBool { // 7.3 这里判断返回值，true则打印，false则不打印
		// 11.1 注释掉一下重复的内容，将重复的内容让repeatInfo函数去做
		//nowTime := getNowTime()
		//// 9.1 报错日志，正常都需要报错的文件名、函数名、行号等信息，所以这里就是提供这些。
		//fileName, funcName, line := callUser(2)
		//// 10.1 将DEBUG(数字类型)，转成字符串类型，返回过来，放入下面fmt.Printf中使用
		//strLogLevel := LogLevelToStr(DEBUG)
		//fmt.Printf("[%s %s] [%s:%s:%v] : %s\n", nowTime, strLogLevel, fileName, funcName, line, msg)

		// 11.2 调用repeatInfo函数，将重复的内容全部交给repeatInfo去做
		//repeatInfo(DEBUG, msg)
		// 12.2 支持msg为格式化的内容传参,且支持格式化的内容中有多个格式化的值("id:%d, name:%s", id ,name 这种方式)
		repeatInfo(DEBUG, formatMsg, a...)
	}
}

// 3.Info 方法，日志打印
//func (con ConsoleLogger) Info(msg string) {
// 12.2 更换格式化传参
func (con ConsoleLogger) Info(formatMsg string, a ...interface{}) {
	if con.CompareLogLevel(INFO) { // 7.4 这里简写判断为true/false
		// 11.2 调用repeatInfo函数，将重复的内容全部交给repeatInfo去做
		//repeatInfo(INFO, msg)
		// 12.3 支持msg为格式化的内容传参,且支持格式化的内容中有多个格式化的值("id:%d, name:%s", id ,name 这种方式)
		repeatInfo(INFO, formatMsg, a...)
	}
}

// 3.Error 方法，日志打印
//func (con ConsoleLogger) Error(msg string) {
func (con ConsoleLogger) Error(formatMsg string, a ...interface{}) {
	if con.CompareLogLevel(ERROR) { // 7.4 这里简写判断为true/false
		// 11.2 调用repeatInfo函数，将重复的内容全部交给repeatInfo去做
		//repeatInfo(ERROR, msg)
		repeatInfo(ERROR, formatMsg, a...)
	}
}

// 3.Fatal 方法，日志打印
//func (con ConsoleLogger) Fatal(msg string) {
func (con ConsoleLogger) Fatal(formatMsg string, a ...interface{}) {
	if con.CompareLogLevel(FATAL) { // 7.4 这里简写判断为true/false
		// 11.2 调用repeatInfo函数，将重复的内容全部交给repeatInfo去做
		//repeatInfo(FATAL, msg)
		repeatInfo(FATAL, formatMsg, a...)
	}
}
