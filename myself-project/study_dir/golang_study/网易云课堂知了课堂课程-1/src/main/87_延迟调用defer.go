package main

import "fmt"

func main8701() {
	// defer语句只能出现在函数内部
	defer fmt.Println("hello") // 最后执行
	defer fmt.Println("老王")    // 其次执行
	fmt.Println("你好")          //最先执行
	//  结果:
	//  你好
	//  老王
	//  hello
}

func test87(v int) {
	res := 2 / v
	fmt.Println(res)
}

func main8702() {
	// defer语句只能出现在函数内部
	defer fmt.Println("hello") // 最后执行
	defer fmt.Println("老王")    // 第三执行
	defer test87(1)            // 第二执行,如果传参是0，则最后报错。
	fmt.Println("你好")          //最先执行
	//  结果:
	//  你好
	//  老王
	//  hello
	//  panic: runtime error: integer divide by zero
}

func main() {
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Println("匿名函数a", a)
		fmt.Println("匿名函数b", b)
	}(a, b)

	a = 100
	b = 200
	fmt.Println("main函数a :", a)
	fmt.Println("main函数b :", b)

	//结果:
	//main函数a : 100
	//main函数b : 200
	//匿名函数a 10
	//匿名函数b 20
}
