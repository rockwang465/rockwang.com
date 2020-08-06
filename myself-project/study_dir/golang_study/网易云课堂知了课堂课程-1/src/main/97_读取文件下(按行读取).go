package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main9701() {
	fp, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer fp.Close()

	// 1.创建文件缓存区
	r := bufio.NewReader(fp)
	// 2.行读取, 换行的标志符为'\n'
	//slice, _ := r.ReadBytes('\n')  // 按换行符读取，读取第一行
	//	//fmt.Println(slice)  // [104 101 108 108 111 32 119 111 114 108 100 32 33 13 10]
	//	//fmt.Println(string(slice))
	//	//// 将ASCII转为字符串
	//	//slice2, _ := r.ReadBytes('\n')  // 按换行符读取，读取第二行
	//	//fmt.Println(string(slice2))

	// 3.for循环按行读取
	for {
		buf, err1 := r.ReadBytes('\n')
		//buf, err1 := r.ReadBytes('h')  // 当然也可以设置其他字符为分隔符进行读取
		// 注意:这里先打印来读取，因为如果读到最后一行，是没有'\n'换行符的，则不会打印了
		fmt.Println(string(buf))
		if err1 != nil {
			if err1 == io.EOF { //读取到结尾了
				break
			} else {
				fmt.Println("读取文件失败")
			}
		}

	}
}

func main() {
	fp, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer fp.Close()
	// 创建文件缓存区
	r := bufio.NewReader(fp)

	for {
		str, err2 := r.ReadString('\n')
		fmt.Println(str)
		if err2 != nil {
			if err2 == io.EOF { //读取到结尾了
				break
			} else {
				fmt.Println("读取文件失败")
			}
		}
	}
}
