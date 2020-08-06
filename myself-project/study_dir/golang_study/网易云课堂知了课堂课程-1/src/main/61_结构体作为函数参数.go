package main

import "fmt"

type Hero struct {
	name  string
	age   int
	power int
}

func test61(h Hero) {
	h.power = 2000
	fmt.Println(h) // {Iron man 30 2000}
}

func main() {
	h := Hero{"Iron man", 30, 1000}
	test61(h)
	fmt.Println(h) // {Iron man 30 1000}
	// 说明结构体作为函数参数是 值传递，所以不影响原值
}
