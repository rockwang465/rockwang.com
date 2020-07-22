package main

import "fmt"

func f2(x, y int) {
	fmt.Println("this is f2 func")
	fmt.Println(x + y)
}

func waiKe(x func(int, int), m, n int) func() {
	// 1.定义个匿名函数，赋值给tmp
	tmp := func() {
		x(m, n)
	}
	// 2.执行tmp()相当于执行x(m,n)
	//tmp() //相当于执行 x(m,n)
	// 3.但这里不是执行tmp()，而是return tmp
	return tmp
}

func main() {
	// 4.那么这里接收返回值，返回值是个函数，即tmp
	returnRes := waiKe(f2, 100, 200)
	// 5.执行这个返回的函数
	// 所以，当你执行returnRes()时，相当于执行f2()函数，而f2就是waiKe()函数中的tmp
	returnRes()
}
