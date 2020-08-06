package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1 只读方式打开文件
	fp, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
		os.Exit(0)
	}
	// 1.2 创建需要读取的长度
	buf := make([]byte, 10) // 创建一个长度为10的字符切片
	// 1.3 读取
	//n, _ := fp.Read(buf) // n代表读取的长度，属于块读取
	//fmt.Println(n)  // 10
	//fmt.Println(string(buf[:n])) // 结果显示了10个字符的内容: gogogo ---
	//n1, _ := fp.Read(buf)
	//fmt.Println(string(buf[:n1]))
	//n2, _ := fp.Read(buf)
	//fmt.Println(string(buf[:n2]))

	// 1.4 for循环读取
	for {
		n, err := fp.Read(buf)
		// io.EOF表示文件的结尾
		// 当读取到文件结尾是，会返回 errors.New("EOF")错误
		if err == io.EOF { // 如果读到文件结尾
			break // 则跳出循环
		}
		fmt.Println((string(buf[:n]))) // 一行一行的打印
	}
}
