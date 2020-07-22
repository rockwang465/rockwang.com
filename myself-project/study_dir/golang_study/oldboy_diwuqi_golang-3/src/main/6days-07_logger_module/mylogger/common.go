package mylogger

import (
	"errors"
	"strings"
	"time"
)

// 8.此处主要用于存放通用的方法（大家都会调用的函数方法），所以下面的代码都是从logger.go中拿过来的
//   由于common.go和logger.go在一个目录下，所以互相之间不需要import，就可以使用互相的变量、函数等。

// 5.1 为下面const中常量定义一个变量类型
type logLevelInt uint16 // 为啥用uint16，我也不知道

// 5.为了用户传入的日志级别需求进行判断，所以这里定义了常量(由于开发环境和生成环境使用的日志级别不同，所以需要进行判断。开发环境一般是DEBUG级别日志，生成环境一般是INFO级别日志)
const (
	// 5.2 int16为DEBUG的数据类型，而iota是数字枚举，DEBUG第一个位置，自动为0，下面的INFO ERROR FATAL自动为 1 ，2 ，3。这样就减少了代码量了
	// 数字形式的日志级别
	UNKNOWN logLevelInt = iota // 老师这里写了一个UNKNOWN，用于下面6.2中判断，在switch的default中返回使用的
	DEBUG
	INFO
	ERROR
	FATAL
)

// 6.判断日志级别，将字符串的级别return为 数字的级别，方便用 > = < 进行判断
func LogLevelToInt(logLevelStr string) (logLevelInt, error) { // Rock注意1: 这里返回值不能写成(logLevelInt, err error)，否则系统会理解成 logLevelInt和err都为error类型。
	// 6.1 先将字符转小写(大写也行)，防止用户输入的大小写不同，不好下面判断
	logLevelStr = strings.ToLower(logLevelStr)

	// 6.2 判断用户传入的字符串的日志级别进来，即 logLevel 形参
	switch logLevelStr {
	case "debug": // 当用户传参进来是debug/DEBUG,则返回一个对应的上面定义的常量int16的值回去，方便用 大于 等于 小于进行判断
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		// 由于用户可能传入错误的日志级别，所以这里需要返回一个error信息给用户，用于提醒
		err := errors.New("输入的为无效的日志级别")
		return UNKNOWN, err
	}
}

// 10 将上面的DEBUG INFO ERROR FATAL数字类型，转为字符串 "Debug Info Error Fatal"
func LogLevelToStr(lvInt logLevelInt) string {
	switch lvInt {
	case DEBUG:
		return "Debug"
	case INFO:
		return "Info"
	case ERROR:
		return "Error"
	case FATAL:
		return "Fatal"
	}
	return "Debug" // 默认返回值返回Debug，其实这里应该是个报错，如果不为上面的，就报错提示。这里懒了，就不多写了
}

// 4.获取当前时间
func getNowTime() (format1, format2 string) {
	NT := time.Now()
	// 注意 "2006-01-02 15:04:05" 这里是固定的，不可以随便写
	//return NT.Format("2006-01-02 15:04:05")
	format1 = NT.Format("2006-01-02 15:04:05") // 打印日志信息时，加的时间戳
	format2 = NT.Format("2006010215040500")    // 切割日志文件，备份源文件时加的时间戳
	return
}
