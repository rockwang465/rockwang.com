package main

import (
	"fmt"
	"strings"
) // 字符串处理函数模块

func main8901() {
	str1 := "hello world !"
	str2 := "e"
	// 1.查找函数: 查找一个字符串在另一个字符串中是否存在 -- Contains
	// Contains，返回bool值--->模糊查找，因为只有true/false
	b := strings.Contains(str1, str2)
	fmt.Println(b) // true
}

func main8902() {
	slice := []string{"123", "456", "789"}

	// 2.拼接切片中的字符串--Join
	str := strings.Join(slice, "")
	fmt.Println(str) // 123456789
	//fmt.Printf("%T\n", str) // string类型
}

func main8903() {
	str1 := "hello world !"
	str2 := "e"

	// 3.查找一个字符串在另一个字符串中的索引值 --Index
	i := strings.Index(str1, str2)
	fmt.Println(i) // 1
}

func main8904() {
	str1 := "我爱生活。"
	// 4.重复字符串
	str2 := strings.Repeat(str1, 5) // 重复次数
	fmt.Println(str2)               // 我爱生活。我爱生活。我爱生活。我爱生活。我爱生活。
}

func main8905() {
	str1 := "性感网友，在线取名，到底是否性感"

	// 5.字符串替换 -- Replace
	str2 := strings.Replace(str1, "性感", "帅气", -1) // "性感"替换为"帅气"，替换次数: -1为替换所有，大于0为替换次数。
	fmt.Println(str2)                             // 帅气网友，在线取名，到底是否帅气
}

func main8906() {
	str1 := "139-2916-0134"

	// 6.字符串分割 -- Split
	str2 := strings.Split(str1, "-") // 按"-"分割
	fmt.Println(str2)                // [139 2916 0134] 返回的是切片
	//fmt.Printf("%T\n", str2)  // []string
}

func main8907() {
	str1 := "    Who are you ?    "
	// 7.去除字符串中的空格，转成【切片】
	str2 := strings.Fields(str1) // 只能按空格分割，所以不及Split好用
	fmt.Println(str2)            // [Who are you ?]，切片中有4个值。 返回值为[]string切片
	//fmt.Printf("%T\n",str2)  // []string
}

func main8908() {
	str1 := "====I==am==boy===="

	// 8.去掉字符串头尾的指定内容 -- Trim
	str2 := strings.Trim(str1, "=") // 去除头尾的"="
	fmt.Println(str2)               // I==am==boy  返回值是字符串
}
