package main

import "fmt"

func test88A() {
	fmt.Println("test88A")
}
func test88B(i int) {
	// 设置recover(), 防止程序崩溃，拦截程序发生的异常
	defer func() {
		// 方法一:  不打印任何信息[普通用法]
		//recover()

		// 方法二: 打印报错，但是当没报错，还是会打印: nil [不好用]
		//fmt.Println(recover())  // 也可以打印异常报错信息

		// 方法三: if判断，解决打印nil的问题 [最佳用法]
		if err := recover(); err != nil { // 如果err!=nil，说明有报错
			fmt.Println(err) // 有报错，则打印报错信息
		}

	}() // ()代表调用匿名函数

	var a [3]int
	a[i] = 999
	//fmt.Println("test88B ",a)
}
func test88C() {
	fmt.Println("test88C")
}

func main() {
	test88A()
	test88B(0) // 当传入值大于2，则报错，会调用recover()
	test88C()
}
