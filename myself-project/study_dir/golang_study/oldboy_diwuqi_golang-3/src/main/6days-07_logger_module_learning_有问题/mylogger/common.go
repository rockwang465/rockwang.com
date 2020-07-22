package mylogger

import (
	"errors"
	"path"
	"runtime"
	"strings"
)

func logLevelToInt(lv string) (int16, error) {
	lv = strings.ToLower(lv)
	err := errors.New("错误的日志等级")
	switch lv {
	case "unknown":
		return 0, nil
	case "debug":
		return 1, nil
	case "info":
		return 2, nil
	case "error":
		return 3, nil
	case "fatal":
		return 4, nil
	}
	return 0, err
}

// 6 添加计算日志报错的文件名、行号、函数名
func callUser(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		panic("Error : runtime.Caller failed")
	}
	funcName = runtime.FuncForPC(pc).Name()
	baseName := path.Base(file)
	splitFileName := strings.Split(file, "/")
	lenFileName := len(splitFileName)
	fileNameDir := splitFileName[lenFileName-2] // 拿到倒数第二个路径名
	fileName = fileNameDir + "/" + baseName     // 拼接文件名，正常报错都是这样的文件名，更规范
	return

}
