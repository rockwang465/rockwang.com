package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	// 定义数组指针(第一指针，指向数组)
	var p *[3]int
	// 指针和数组建立关系
	p = &arr
	fmt.Println(*p)

	// 通过指针操作数组的方法:
	(*p)[0] = 111 // 方法一: 必须要括号把*p括起来才行
	p[1] = 222    // 方法二: 直接拿p去操作
	fmt.Println(*p)
}
