package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

const skip  = 1

func errorPrint() {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("Error : error Print")
		return
	}
	fileName := path.Base(file)
	lenFuncName := len(strings.Split(runtime.FuncForPC(pc).Name(), "."))        // 先取长度
	funcName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[lenFuncName-1] // 从结果main.main中，取后面一个函数名就好，不然太多有点难看
	fmt.Printf("func name:%v\n", funcName)                                      // 正常应该是 main.main
	fmt.Printf("file name:%v\n", fileName)                                      // test_runtime_Caller.go
	fmt.Printf("line:%v\n", line)                                               // 27
	errMsg := "init etcd failed"
	fmt.Printf("Error: func_name: %v, file_name:[%v], line:[%v] :%s ", pc, fileName, line, errMsg)
}

func main() {
	errorPrint()
}
