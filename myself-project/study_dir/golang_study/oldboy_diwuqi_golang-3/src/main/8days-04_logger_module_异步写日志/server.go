package main

import (
	"./mylogger"
	"time"
)

var log mylogger.Logger // 声明一个全局的接口变量 (多态方式?)
//var logC mylogger.Logger // 感觉也可以声明多个全局变量来同时调用2中方式的log功能(console和file)

func main() {
	// NewLog(日志级别) // 传入当前开发环境/生成环境使用的。开发环境一般是DEBUG级别日志，生成环境一般是INFO级别日志
	// ------- 普通方式 -------
	// A. 向屏幕输入日志
	//log := mylogger.NewConsoleLog("debug")
	// B. 向文件输入日志
	//log := mylogger.NewFileLog("debug", "rock.log", "./", 1*1024*1024)

	// ------- 较为高级的方式: 多态 -------
	// 1024 = 1k, 10*1024*1024 = 10mb
	//log = mylogger.NewConsoleLog("debug")                              //向终端输出日志
	log = mylogger.NewFileLog("debug", "rock.log", "./", 10*1024*1024) // 向文件输出日志

	// ------- 最高级方式: 封装上面两种到一起 -------
	// log = mylogger.NewLogger("console")  //  通过这一个函数的传参，来判断是使用 NewConsoleLog 或  NewFileLog，功能更强大

	for {
		log.Debug("Welcome use myLogger module")
		//time.Sleep(time.Second * 2)
		log.Info("Server is running")
		//time.Sleep(time.Second * 2)
		id := 10086
		name := "Rock"
		// 12.4 支持格式化字变量传参
		log.Error("Not found /etc/nginx/conf.d directory, error account: id:%d, user:%s", id, name)
		//time.Sleep(time.Second * 2)
		log.Fatal("Server is down , bye bye ...")
		//time.Sleep(time.Second * 2)
		log.Info("Restart Server")
		//time.Sleep(time.Second * 4)
		time.Sleep(time.Millisecond * 100)
	}
}

// 总的日志代码需求
// 1. 日志分级别:
//   Debug Trace Info Warning Error Fatal
// 2. 日志要支持开关，比如开发的时候输出Debug级别日志，上线后输出Info级别日志
// 3. 完整的日志记录内容要包含: 时间、行号、文件名、日志级别、日志信息
// 4. 支持格式化传入 需要打印的日志信息(如: "id:%s name:%s", 2, rock 这种方式传入日志内容)
// 5. 支持往不同的地方输出: 屏幕 和 文件
// 6. 日志文件要切割
//   日志切割逻辑: 关闭当前打开的日志文件，
//                重命名为bak+时间，
//                打开新文件，
//                将当前打开的文件的对象赋值给结构体之前保存的文件对象。这样后面则使用新的文件对象进行操作了。
//   A. 按文件大小切割： 在记录日志前判断是否达到预定义文件大小，超过则切割
//   B. 按时间小时切割:  在日志结构体设置一个字段，记录上一次切割的小时数
//                      在写日志之前检查一下当前的小时数和之前保持的是否一致，不一致就切割
// 7. 支持interface多态，可以同时输出到屏幕 及 日志文件中
// 8. 在第8天的课程里要求: 改为异步写日志，需要用到channel和select


// 最后Rock总结一下坑：
// 1.结构体传输，例如新增个函数，当做结构体的class时，请一定要加* ，一定要加 *，真的一定要加* ，必须指针传输，不然GG~