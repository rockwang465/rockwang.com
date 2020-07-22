package main

import "fmt"

func demo1(m map[int]string) {
	m[4] = "甄姬"
	fmt.Println(m)        // map[1:小乔 2:貂蝉 3:大乔 4:甄姬]
	fmt.Printf("%p\n", m) // 0xc000068330 内存地址相同
}

func main() {
	m := map[int]string{1: "小乔", 2: "貂蝉", 3: "大乔"}
	// 地址传递，所以demo1中修改map字典的值会影响当前的map字典
	demo1(m)
	fmt.Println(m)        // map[1:小乔 2:貂蝉 3:大乔 4:甄姬]
	fmt.Printf("%p\n", m) // 0xc000068330 内存地址相同
}
