package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]byte, 0, 1024) // 创建字符切片，长度0，容量1024

	// 1.将其他类型转换成字符串，添加到字符切片中。
	// 第一种: 将布尔类型添加到字符切片中
	slice = strconv.AppendBool(slice, false)
	//fmt.Println(slice)  // [102 97 108 115 101]
	//fmt.Println(string(slice))  // false

	// 第二种: 整数型添加到字符切片中
	slice = strconv.AppendInt(slice, 4, 2) // 4为整型值，2为进制
	//fmt.Println(string(slice))  // false100

	// 第三种: 浮点型添加到字符切片中
	slice = strconv.AppendFloat(slice, 3.141592653, 'f', 4, 64)
	// 第四种: 字符串型添加到字符切片中
	slice = strconv.AppendQuote(slice, "hello")
	fmt.Println(string(slice)) // false1003.1416"hello"
}
