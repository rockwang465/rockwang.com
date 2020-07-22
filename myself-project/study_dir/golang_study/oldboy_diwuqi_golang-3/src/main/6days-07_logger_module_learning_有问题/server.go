package main

import (
	"./mylogger"
	"time"
)

func main() {
	name := "rock"
	id := 666
	for {
		// a.往屏幕打印日志
		//log := mylogger.NewConsoleLog("debug")

		// b.往文件输出日志
		log := mylogger.NewFileLog("debug", "nginx.log", "nginx_err.log","./", 1*1024*1024)
		log.Debug("Welcome use Nginx server")
		log.Info("Server is running")
		log.Info("id: %d, name: %s , Welcome back ...", id, name)
		log.Error("Not found /etc/nginx/conf.d/ directory or file")
		time.Sleep(time.Millisecond * 1)
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

// 最后Rock总结一下坑：
// 1.结构体传输，例如新增个函数，当做结构体的class时，请一定要加* ，一定要加 *，真的一定要加* ，必须指针传输，不然GG~
