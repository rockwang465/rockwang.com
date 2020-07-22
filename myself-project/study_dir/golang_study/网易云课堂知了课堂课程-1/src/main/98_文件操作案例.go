package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var srcFileName string
	var dstFileName string

	// 1.输入文件名
	fmt.Println("请输入源文件名: ")
	fmt.Scan(&srcFileName)
	fmt.Println("请输入目的文件名: ")
	fmt.Scan(&dstFileName)

	// 2.确认没有重名
	if srcFileName == dstFileName {
		fmt.Println("Error : 源文件名与目标文件名相同，请重新输入！")
		return
	} else {
		fmt.Println("Info : 输入文件名正常。")
	}

	// 3.打开和创建两个文件
	sf, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("Error : 源文件打开失败", err1)
		return
	}

	df, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("Error : 目标文件创建失败", err2)
	}

	// 4.关闭两个文件
	defer sf.Close()
	defer df.Close()

	// 5.操作文件
	buf := make([]byte, 1024*4) // 定义每次块读取的量为 4KB
	//buf := make([]byte, 10)  // 定义每次 10 字节
	for {
		// 5.1读取打开的源文件 -- 按块读取
		res1, err3 := sf.Read(buf)
		if err3 != nil { // 如果读取有报错
			if err3 == io.EOF { // 如果读到结尾，则结束
				break
			} else { // 其他异常则报错
				fmt.Println("Error : 文件按块读取失败", err3)
				return
			}
		} else { // 5.2写入文件的
			//fmt.Println(res1)  // 返回读取到的字节数
			//fmt.Println(string(buf[:res1]))
			_, err4 := df.Write(buf[:res1])
			if err4 != nil {
				fmt.Println("Error : 文件写入失败", err4)
				return
			}
		}
	}

}
