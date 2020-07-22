package main

import "fmt"

//func Swap63(a, b int) {
//	a, b = b, a
//	fmt.Printf("Swap函数: a=%d,b=%d\n", a, b)
//}

func Swap64(a, b *int) {
	//a, b = b, a  // 原值不会变化
	*a, *b = *b, *a // 原值会变化
	fmt.Printf("Swap函数: a=%d,b=%d\n", *a, *b)
}

func main() {
	a := 10
	b := 20
	//Swap63(a, b)  // 值传递不会变化
	Swap64(&a, &b) // 地址传递
	fmt.Printf("main函数: a=%d,b=%d", a, b)
}
