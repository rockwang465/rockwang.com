package main

import (
	"fmt"
	"strconv"
)

func main9101() {
	// 1.字符串转为字符切片,且以ASCII形式显示  -- []byte()
	str1 := "hello world !"
	str2 := []byte(str1)
	fmt.Println(str2) // [104 101 108 108 111 32 119 111 114 108 100 32 33]
}

func main9102() {
	// 2.字符切片转字符串, 且将ASCII形式的转为字符 -- string()
	slice := []byte{'h', 'e', 'l', 'l', 97, 'o'}
	str := string(slice)
	fmt.Println(str) // hellao
}

func main9103() {
	// 3.将其类型转为字符串类型 -- Format
	// 第一种: 布尔类型转换 -- FormatBool
	b := false
	str1 := strconv.FormatBool(b)
	fmt.Println(str1)        // false
	fmt.Printf("%T\n", str1) // string

	// 第二种: 进制转换 -- FormatInt
	// strconv.FormatInt(换算进制的值, 换算的进制)
	str2 := strconv.FormatInt(120, 2) // 计算机中有2-36进制
	fmt.Println(str2)                 // 1111000

	// 第三种: 浮点型转换 -- FormatFloat
	// strconv.FormatFloat(浮点数值, '浮点类型', 保留小数位, 浮点类型中的格式)
	str3 := strconv.FormatFloat(3.141592653, 'f', 4, 64)
	fmt.Println(str3) // 3.1416

	// 第四种: int整型转换 -- Itoa
	str4 := strconv.Itoa(123)
	fmt.Println(str4)
}

func main() {
	// 4.字符串转成其他类型  -- Parse
	// 第一种:  字符串转布尔类型
	v1, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("类型转换出错")
	} else {
		fmt.Println(v1)        // true
		fmt.Printf("%T\n", v1) // bool
	}

	// 第二种: 字符串转整型
	// 接收值, err := strconv.ParseInt("字符串", 进制, 整型类型)
	// 注意:如果err值也接收，则可以显示报错内容。下面没有接收err值。
	v2, _ := strconv.ParseInt("abc", 10, 64)
	// 这里字符串是abc，是无法转换成10进制的，所以返回0.
	fmt.Println(v2) // 0

	// 第三种: 字符串转浮点型
	v3, _ := strconv.ParseFloat("3.141592653", 64)
	fmt.Println(v3) // 3.141592653

	// 第四种: 字符串转整型
	v4, _ := strconv.Atoi("123")
	fmt.Println(v4) // 123
}
