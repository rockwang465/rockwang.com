package main

import "fmt"

func demo(s []int) {
	fmt.Println("demo函数第一次打印:", s)
	// 1.2 修改参数
	//s[0] = 666  // 1.3切片传参，由于是内存地址传递，所以修改值会影响原来的切片

	// 2.1 append追加值，不会影响原来的切片
	s = append(s, 66, 77, 88, 99)
	fmt.Println("demo函数第二次打印:", s)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	// 1.1 传参
	demo(s)
	fmt.Println("main函数中最后打印:", s)
}
