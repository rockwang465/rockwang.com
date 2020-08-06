package main

import (
	"fmt"
	"os"
)

func main() {
	// 1.打开文件
	// 1.1 os.Open为只读打开，不支持写操作
	//fp, err := os.Open("a.txt")

	// 1.2 os.OpenFile支持读写操作
	// 文件操作权限为: O_RDWR读写模式; 保存后文件的权限为: 6可读可写。
	fp, err := os.OpenFile("a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("文件打开失败")
	} else {
		fmt.Println(*fp)
	}

	// 1.3 然后进行写操作
	fp.WriteString("hello")
	fp.WriteAt([]byte("hello"), 25)

	defer fp.Close()
}
